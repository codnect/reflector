package reflector

import (
	"errors"
	"reflect"
)

type Boolean interface {
	Type
	Instantiable
	CanSet() bool
	Value() (bool, error)
	SetValue(val bool) error
}

type booleanType struct {
	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (b *booleanType) Name() string {
	return b.reflectType.Name()
}

func (b *booleanType) PackageName() string {
	return ""
}

func (b *booleanType) PackagePath() string {
	return ""
}

func (b *booleanType) HasValue() bool {
	return b.reflectValue != nil
}

func (b *booleanType) Parent() Type {
	return b.parent
}

func (b *booleanType) ReflectType() reflect.Type {
	return b.reflectType
}

func (b *booleanType) ReflectValue() *reflect.Value {
	return b.reflectValue
}

func (b *booleanType) Compare(another Type) bool {
	return false
}

func (b *booleanType) CanSet() bool {
	if b.reflectValue == nil {
		return false
	}

	return b.reflectValue.CanSet()
}

func (b *booleanType) Value() (bool, error) {
	if b.reflectValue == nil {
		return false, errors.New("value reference is nil")
	}

	return b.reflectValue.Interface().(bool), nil
}

func (b *booleanType) SetValue(val bool) error {
	if !b.CanSet() {
		return errors.New("value cannot be set")
	}

	b.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (b *booleanType) Instantiate() Value {
	return &value{
		reflect.New(b.reflectType),
	}
}
