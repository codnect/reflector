package reflector

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Function interface {
	Type
	IsExported() bool
	Parameters() []Type
	NumParameter() int
	Results() []Type
	NumResult() int
	IsVariadic() bool
	Invoke(args ...any) ([]any, error)
}

type functionType struct {
	name       string
	pkgPath    string
	isExported bool

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (f *functionType) Name() string {
	if f.name != "" {
		return f.name
	}

	if f.reflectValue == nil {
		return f.getFunctionName()
	}

	name := runtime.FuncForPC(f.reflectValue.Pointer()).Name()
	dotLastIndex := strings.LastIndex(name, ".")

	if dotLastIndex != -1 {
		return name[dotLastIndex+1:]
	}

	if f.name == "" {
		return f.getFunctionName()
	}

	return name
}

func (f *functionType) getFunctionName() string {
	var builder strings.Builder
	builder.WriteString("func(")

	for index, parameter := range f.Parameters() {
		builder.WriteString(parameter.Name())
		if index != f.NumParameter()-1 {
			builder.WriteString(",")
		}
	}

	builder.WriteString(")")

	if f.NumResult() > 1 {
		builder.WriteString(" (")
	}

	for index, result := range f.Results() {
		builder.WriteString(result.Name())
		if index != f.NumResult()-1 {
			builder.WriteString(",")
		}
	}

	if f.NumResult() > 1 {
		builder.WriteString(")")
	}

	return builder.String()
}

func (f *functionType) PackageName() string {
	if f.pkgPath != "" {
		name := f.pkgPath
		slashLastIndex := strings.LastIndex(name, "/")

		if slashLastIndex != -1 {
			name = name[slashLastIndex+1:]
		}

		return name
	}

	if f.reflectValue == nil {
		return ""
	}

	name := runtime.FuncForPC(f.reflectValue.Pointer()).Name()
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

func (f *functionType) PackagePath() string {
	if f.pkgPath != "" {
		return f.pkgPath
	}

	if f.reflectValue == nil {
		return ""
	}

	name := runtime.FuncForPC(f.reflectValue.Pointer()).Name()
	dotLastIndex := strings.LastIndex(name, ".")

	if dotLastIndex != -1 {
		name = name[:dotLastIndex]
	}

	dotLastIndex = strings.LastIndex(name, ".")

	if dotLastIndex != -1 {
		name = name[:dotLastIndex]
	}

	return name
}

func (f *functionType) CanSet() bool {
	if f.reflectValue == nil {
		return false
	}

	return f.reflectValue.CanSet()
}

func (f *functionType) HasValue() bool {
	return f.reflectValue != nil
}

func (f *functionType) Value() (any, error) {
	if f.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return f.reflectValue.Interface(), nil
}

func (f *functionType) SetValue(val any) error {
	if !f.CanSet() {
		return errors.New("value cannot be set")
	}

	f.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (f *functionType) Parent() Type {
	return f.parent
}

func (f *functionType) ReflectType() reflect.Type {
	return f.reflectType
}

func (f *functionType) ReflectValue() *reflect.Value {
	return f.reflectValue
}

func (f *functionType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return f.reflectType == another.ReflectType()
}

func (f *functionType) IsInstantiable() bool {
	return false
}

func (f *functionType) Instantiate() (Value, error) {
	return nil, errors.New("functions are not instantiable")
}

func (f *functionType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if f.reflectValue == nil {
		return f.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return f.reflectValue.CanConvert(typ.ReflectType())
}

func (f *functionType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if f.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !f.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := f.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (f *functionType) IsExported() bool {
	return f.isExported
}

func (f *functionType) Parameters() []Type {
	parameters := make([]Type, 0)
	numIn := f.reflectType.NumIn()

	for index := 0; index < numIn; index++ {
		typ := f.reflectType.In(index)
		parameters = append(parameters, typeOf(nil, typ, nil, nil))
	}
	return parameters
}

func (f *functionType) NumParameter() int {
	return f.reflectType.NumIn()
}

func (f *functionType) Results() []Type {
	results := make([]Type, 0)
	numOut := f.reflectType.NumOut()

	for index := 0; index < numOut; index++ {
		typ := f.reflectType.Out(index)
		results = append(results, typeOf(nil, typ, nil, nil))
	}
	return results
}

func (f *functionType) NumResult() int {
	return f.reflectType.NumOut()
}

func (f *functionType) IsVariadic() bool {
	return f.reflectType.IsVariadic()
}

func (f *functionType) Invoke(args ...any) ([]any, error) {
	if f.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if (f.IsVariadic() && len(args) < f.NumParameter()) || (!f.IsVariadic() && len(args) != f.NumParameter()) {
		return nil, fmt.Errorf("invalid parameter count, expected %d but got %d", f.NumParameter(), len(args))
	}

	inputs := make([]reflect.Value, 0)

	var variadicType Slice

	if f.IsVariadic() {
		paramType := f.Parameters()[f.NumParameter()-1]
		variadicType = ToSlice(paramType)
	}

	for index, arg := range args {
		actualParamType := TypeOfAny(arg)

		if f.IsVariadic() && index > f.NumResult() {
			if arg == nil {
				inputs = append(inputs, reflect.New(variadicType.Elem().ReflectType()).Elem())
				continue
			} else if !actualParamType.CanConvert(variadicType.Elem()) {
				return nil, fmt.Errorf("expected %s but got %s at index %d", variadicType.Elem().Name(), actualParamType.Name(), index)
			}

			inputs = append(inputs, reflect.ValueOf(arg))
			continue
		}

		expectedParamType := f.Parameters()[index]

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
	results := f.reflectValue.Call(inputs)

	for _, outputParam := range results {
		outputs = append(outputs, outputParam.Interface())
	}

	return outputs, nil
}
