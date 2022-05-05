package reflector

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
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
	//ReflectMethod() reflect.Method
}

type methodType struct {
	name       string
	pkgPath    string
	isExported bool
	receiver   Type

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (m *methodType) Name() string {
	if m.name != "" {
		return m.name
	}

	if m.reflectValue == nil {
		return m.getFunctionName()
	}

	name := runtime.FuncForPC(m.reflectValue.Pointer()).Name()
	dotLastIndex := strings.LastIndex(name, ".")

	if dotLastIndex != -1 {
		return name[dotLastIndex+1:]
	}

	if m.name == "" {
		return m.getFunctionName()
	}

	return name
}

func (m *methodType) getFunctionName() string {
	var builder strings.Builder
	builder.WriteString("func(")

	for index, parameter := range m.Parameters() {
		builder.WriteString(parameter.Name())
		if index != m.NumParameter()-1 {
			builder.WriteString(",")
		}
	}

	builder.WriteString(")")

	if m.NumResult() > 1 {
		builder.WriteString(" (")
	}

	for index, result := range m.Results() {
		builder.WriteString(result.Name())
		if index != m.NumResult()-1 {
			builder.WriteString(",")
		}
	}

	if m.NumResult() > 1 {
		builder.WriteString(")")
	}

	return builder.String()
}

func (m *methodType) PackageName() string {
	if m.pkgPath != "" {
		name := m.pkgPath
		slashLastIndex := strings.LastIndex(name, "/")

		if slashLastIndex != -1 {
			name = name[slashLastIndex+1:]
		}

		return name
	}

	if m.reflectValue == nil {
		return ""
	}

	name := runtime.FuncForPC(m.reflectValue.Pointer()).Name()
	dotLastIndex := strings.LastIndex(name, ".")

	if dotLastIndex != -1 {
		name = name[:dotLastIndex]
	}

	slashLastIndex := strings.LastIndex(name, "/")

	if slashLastIndex != -1 {
		name = name[slashLastIndex+1:]
	}

	return name
}

func (m *methodType) PackagePath() string {
	return ""
}

func (m *methodType) CanSet() bool {
	if m.reflectValue == nil {
		return false
	}

	return m.reflectValue.CanSet()
}

func (m *methodType) HasValue() bool {
	return m.reflectValue != nil
}

func (m *methodType) Value() (any, error) {
	if m.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return m.reflectValue.Interface(), nil
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
	return m.reflectType
}

func (m *methodType) ReflectValue() *reflect.Value {
	return m.reflectValue
}

func (m *methodType) Compare(another Type) bool {
	return false
}

func (m *methodType) IsInstantiable() bool {
	return false
}

func (m *methodType) Instantiate() (Value, error) {
	return nil, nil
}

func (m *methodType) IsExported() bool {
	return m.isExported
}

func (m *methodType) Receiver() Type {
	return nil
}

func (m *methodType) Parameters() []Type {
	parameters := make([]Type, 0)
	numIn := m.reflectType.NumIn()

	for index := 0; index < numIn; index++ {
		typ := m.reflectType.In(index)
		parameters = append(parameters, typeOf(nil, typ, nil, nil))
	}
	return parameters
}

func (m *methodType) NumParameter() int {
	return m.reflectType.NumIn()
}

func (m *methodType) Results() []Type {
	results := make([]Type, 0)
	numOut := m.reflectType.NumOut()

	for index := 0; index < numOut; index++ {
		typ := m.reflectType.Out(index)
		results = append(results, typeOf(nil, typ, nil, nil))
	}
	return results
}

func (m *methodType) NumResult() int {
	return m.reflectType.NumOut()
}

func (m *methodType) IsVariadic() bool {
	return m.reflectType.IsVariadic()
}

func (m *methodType) Invoke(args ...any) ([]any, error) {
	if m.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if (m.IsVariadic() && len(args) < m.NumParameter()) || (!m.IsVariadic() && len(args) != m.NumParameter()) {
		return nil, fmt.Errorf("invalid parameter count, expected %d but got %d", m.NumParameter(), len(args))
	}

	inputs := make([]reflect.Value, 0)

	var variadicType Slice

	if m.IsVariadic() {
		paramType := m.Parameters()[m.NumParameter()-1]
		variadicType, _ = ToSlice(paramType)
	}

	for index, arg := range args {
		actualParamType := TypeOfAny(arg)

		if m.IsVariadic() && index > m.NumResult() {
			if arg == nil {
				inputs = append(inputs, reflect.New(variadicType.Elem().ReflectType()).Elem())
				continue
			} else if variadicType.Elem().Name() != "any" && actualParamType.Name() != variadicType.Elem().Name() {
				return nil, fmt.Errorf("expected %s but got %s at index %d", variadicType.Elem().Name(), actualParamType.Name(), index)
			}

			inputs = append(inputs, reflect.ValueOf(arg))
			continue
		}

		expectedParamType := m.Parameters()[index]

		if arg == nil {
			inputs = append(inputs, reflect.New(expectedParamType.ReflectType()).Elem())
		} else {
			if expectedParamType.Name() != "any" && actualParamType.Name() != expectedParamType.Name() {
				return nil, fmt.Errorf("expected %s but got %s at index %d", expectedParamType.Name(), actualParamType.Name(), index)
			}
			inputs = append(inputs, reflect.ValueOf(arg))
		}
	}

	outputs := make([]any, 0)
	results := m.reflectValue.Call(inputs)

	for _, outputParam := range results {
		outputs = append(outputs, outputParam.Interface())
	}

	return outputs, nil
}
