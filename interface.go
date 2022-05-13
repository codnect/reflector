package reflector

import (
	"errors"
	"reflect"
	"strings"
)

type Interface interface {
	Type
	Underlying() Type
	Methods() []Method
	NumMethod() int
}

type interfaceType struct {
	parent         Type
	underlyingType Type
	reflectType    reflect.Type
	reflectValue   *reflect.Value
}

func (i *interfaceType) Name() string {
	if i.reflectType.Name() == "" {
		return "any"
	}

	return i.reflectType.Name()
}

func (i *interfaceType) PackageName() string {
	name := i.reflectType.PkgPath()
	slashLastIndex := strings.LastIndex(name, "/")

	if slashLastIndex != -1 {
		name = name[slashLastIndex+1:]
	}

	return name
}

func (i *interfaceType) PackagePath() string {
	return i.reflectType.PkgPath()
}

func (i *interfaceType) CanSet() bool {
	if i.reflectValue == nil {
		return false
	}

	return i.reflectValue.CanSet()
}

func (i *interfaceType) HasValue() bool {
	return i.reflectValue != nil
}

func (i *interfaceType) Value() (any, error) {
	if i.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return i.reflectValue.Interface(), nil
}

func (i *interfaceType) SetValue(val any) error {
	if !i.CanSet() {
		return errors.New("value cannot be set")
	}

	i.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (i *interfaceType) Parent() Type {
	return i.parent
}

func (i *interfaceType) ReflectType() reflect.Type {
	return i.reflectType
}

func (i *interfaceType) ReflectValue() *reflect.Value {
	return i.reflectValue
}

func (i *interfaceType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return i.reflectType == another.ReflectType()
}

func (i *interfaceType) IsInstantiable() bool {
	return false
}

func (i *interfaceType) Instantiate() (Value, error) {
	return nil, errors.New("interfaces are not instantiable")
}

func (i *interfaceType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if i.reflectValue == nil {
		return i.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return i.reflectValue.CanConvert(typ.ReflectType())
}

func (i *interfaceType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if i.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !i.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := i.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (i *interfaceType) Underlying() Type {
	return i.underlyingType
}

func (i *interfaceType) Methods() []Method {
	functions := make([]Method, 0)
	numMethod := i.reflectType.NumMethod()

	for index := 0; index < numMethod; index++ {
		function := i.reflectType.Method(index)

		method := &methodType{
			parent:        i,
			reflectMethod: function,
		}

		if i.reflectValue != nil {
			underlyingMethod, _ := i.underlyingType.ReflectType().MethodByName(function.Name)
			method.underlyingMethod = &underlyingMethod
		}

		functions = append(functions, method)
	}

	return functions
}

func (i *interfaceType) NumMethod() int {
	return i.reflectType.NumMethod()
}
