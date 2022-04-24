package reflector

import (
	"errors"
	"reflect"
)

type String interface {
	Type
	Instantiable
	CanSet() bool
	Value() (string, error)
	SetValue(val string) error
}

type stringType struct {
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *stringType) Name() string {
	return s.reflectType.Name()
}

func (s *stringType) PackageName() string {
	return ""
}

func (s *stringType) HasValue() bool {
	return s.reflectValue != nil
}

func (s *stringType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *stringType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *stringType) CanSet() bool {
	if s.reflectValue == nil {
		return false
	}

	return s.reflectValue.CanSet()
}

func (s *stringType) Value() (string, error) {
	if s.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return s.reflectValue.Interface().(string), nil
}

func (s *stringType) SetValue(val string) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	s.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (s *stringType) Instantiate() Value {
	return &value{
		reflect.New(s.reflectType),
	}
}
