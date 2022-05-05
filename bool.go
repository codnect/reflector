package reflector

import (
	"errors"
	"reflect"
)

type Boolean interface {
	Type
	BooleanValue() (bool, error)
	SetBooleanValue(val bool) error
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

func (b *booleanType) CanSet() bool {
	if b.reflectValue == nil {
		return false
	}

	return b.reflectValue.CanSet()
}

func (b *booleanType) HasValue() bool {
	return b.reflectValue != nil
}

func (b *booleanType) Value() (any, error) {
	if b.reflectValue == nil {
		return false, errors.New("value reference is nil")
	}

	return b.reflectValue.Interface(), nil
}

func (b *booleanType) SetValue(val any) error {
	if !b.CanSet() {
		return errors.New("value cannot be set")
	}

	switch typedVal := val.(type) {
	case bool:
		return b.SetBooleanValue(typedVal)
	default:
		return errors.New("type is not valid")
	}
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
	if another == nil {
		return false
	}

	return b.reflectType == another.ReflectType()
}

func (b *booleanType) IsInstantiable() bool {
	return true
}

func (b *booleanType) Instantiate() (Value, error) {
	return &value{
		reflect.New(b.reflectType),
	}, nil
}

func (b *booleanType) BooleanValue() (bool, error) {
	if b.reflectValue == nil {
		return false, errors.New("value reference is nil")
	}

	return b.reflectValue.Interface().(bool), nil
}

func (b *booleanType) SetBooleanValue(val bool) error {
	if !b.CanSet() {
		return errors.New("value cannot be set")
	}

	b.reflectValue.Set(reflect.ValueOf(val))
	return nil
}
