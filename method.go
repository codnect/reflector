package reflector

import (
	"errors"
	"fmt"
	"reflect"
)

type Method interface {
	Type
	IsExported() bool
	Receiver() Type
	Parameters() []Type
	NumParameter() int
	Results() []Type
	NumResult() int
	IsVariadic() bool
	Invoke(args ...any) ([]any, error)
	ReflectMethod() reflect.Method
}

type methodType struct {
	parent           Type
	reflectValue     *reflect.Value
	reflectMethod    reflect.Method
	underlyingMethod *reflect.Method
}

func (m *methodType) Name() string {
	return m.reflectMethod.Name
}

func (m *methodType) PackageName() string {
	return m.Parent().PackageName()
}

func (m *methodType) PackagePath() string {
	return m.Parent().PackagePath()
}

func (m *methodType) CanSet() bool {
	return m.ReflectValue().CanSet()
}

func (m *methodType) HasValue() bool {
	return true
}

func (m *methodType) Value() (any, error) {
	return m.ReflectValue().Interface(), nil
}

func (m *methodType) SetValue(val any) error {
	if !m.CanSet() {
		return errors.New("value cannot be set")
	}

	m.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (m *methodType) Parent() Type {
	return m.parent
}

func (m *methodType) ReflectType() reflect.Type {
	return m.reflectMethod.Type
}

func (m *methodType) ReflectValue() *reflect.Value {
	return &m.reflectMethod.Func
}

func (m *methodType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return m.reflectMethod.Type == another.ReflectType()
}
func (m *methodType) IsInstantiable() bool {
	return false
}

func (m *methodType) Instantiate() (Value, error) {
	return nil, errors.New("methods are not instantiable")
}

func (m *methodType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if m.reflectValue == nil {
		return m.ReflectType().ConvertibleTo(typ.ReflectType())
	}

	return m.reflectValue.CanConvert(typ.ReflectType())
}

func (m *methodType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if m.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	val := m.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (m *methodType) IsExported() bool {
	return m.reflectMethod.IsExported()
}

func (m *methodType) Receiver() Type {
	return nil
}

func (m *methodType) Parameters() []Type {
	parameters := make([]Type, 0)
	numIn := m.ReflectType().NumIn()

	index := 1
	if IsInterface(m.parent) {
		index = 0
	}

	for ; index < numIn; index++ {
		typ := m.ReflectType().In(index)
		parameters = append(parameters, typeOf(nil, typ, nil, nil))
	}
	return parameters
}

func (m *methodType) NumParameter() int {
	if IsInterface(m.parent) {
		return m.ReflectType().NumIn()
	}

	return m.ReflectType().NumIn() - 1
}

func (m *methodType) Results() []Type {
	results := make([]Type, 0)
	numOut := m.ReflectType().NumOut()

	for index := 0; index < numOut; index++ {
		typ := m.ReflectType().Out(index)
		results = append(results, typeOf(nil, typ, nil, nil))
	}
	return results
}

func (m *methodType) NumResult() int {
	return m.ReflectType().NumOut()
}

func (m *methodType) IsVariadic() bool {
	return m.ReflectType().IsVariadic()
}

func (m *methodType) Invoke(args ...any) ([]any, error) {
	parent := m.Parent()
	reflectValue := parent.ReflectValue()

	if reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if parent.Parent() != nil {
		reflectValue = parent.Parent().ReflectValue()
	}

	if reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if (m.IsVariadic() && len(args) < m.NumParameter()) || (!m.IsVariadic() && len(args) != m.NumParameter()) {
		return nil, fmt.Errorf("invalid parameter count, expected %d but got %d", m.NumParameter(), len(args))
	}

	inputs := make([]reflect.Value, 0)

	var variadicType Slice

	if m.IsVariadic() {
		paramType := m.Parameters()[m.NumParameter()-1]
		variadicType = ToSlice(paramType)
	}

	for index, arg := range args {
		actualParamType := TypeOfAny(arg)

		if m.IsVariadic() && index > m.NumResult() {
			if arg == nil {
				inputs = append(inputs, reflect.New(variadicType.Elem().ReflectType()).Elem())
				continue
			} else if !actualParamType.CanConvert(variadicType.Elem()) {
				return nil, fmt.Errorf("expected %s but got %s at index %d", variadicType.Elem().Name(), actualParamType.Name(), index)
			}

			inputs = append(inputs, reflect.ValueOf(arg))
			continue
		}

		expectedParamType := m.Parameters()[index]

		if arg == nil {
			inputs = append(inputs, reflect.New(expectedParamType.ReflectType()).Elem())
		} else {
			if !actualParamType.CanConvert(expectedParamType) {
				return nil, fmt.Errorf("expected %s but got %s at index %d", expectedParamType.Name(), actualParamType.Name(), index)
			}
			inputs = append(inputs, reflect.ValueOf(arg))
		}
	}

	outputs := make([]any, 0)

	if parent.Parent() == nil {
		var pointer reflect.Value
		if IsInterface(parent) {
			inputs = append([]reflect.Value{*reflectValue}, inputs...)
		} else {
			pointer = reflect.New(parent.ReflectType())
			pointer.Elem().Set(*reflectValue)
			inputs = append([]reflect.Value{pointer}, inputs...)
		}

	} else {
		inputs = append([]reflect.Value{*reflectValue}, inputs...)
	}

	var results []reflect.Value

	if m.underlyingMethod != nil {
		results = m.underlyingMethod.Func.Call(inputs)
	} else {
		results = m.reflectMethod.Func.Call(inputs)
	}

	for _, outputParam := range results {
		outputs = append(outputs, outputParam.Interface())
	}

	return outputs, nil
}

func (m *methodType) ReflectMethod() reflect.Method {
	return m.reflectMethod
}
