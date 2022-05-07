package reflector

import (
	"errors"
	"reflect"
)

type String interface {
	Type
	StringValue() (string, error)
	SetStringValue(val string) error
}

type stringType struct {
	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *stringType) Name() string {
	return s.reflectType.Name()
}

func (s *stringType) PackageName() string {
	return ""
}

func (s *stringType) PackagePath() string {
	return ""
}

func (s *stringType) CanSet() bool {
	if s.reflectValue == nil {
		return false
	}

	return s.reflectValue.CanSet()
}

func (s *stringType) HasValue() bool {
	return s.reflectValue != nil
}

func (s *stringType) Value() (any, error) {
	if s.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return s.reflectValue.Interface(), nil
}

func (s *stringType) SetValue(val any) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	switch typedVal := val.(type) {
	case string:
		return s.SetStringValue(typedVal)
	default:
		return errors.New("type is not valid")
	}
}

func (s *stringType) Parent() Type {
	return s.parent
}

func (s *stringType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *stringType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *stringType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return s.reflectType == another.ReflectType()
}

func (s *stringType) IsInstantiable() bool {
	return true
}

func (s *stringType) Instantiate() (Value, error) {
	return &value{
		reflect.New(s.reflectType),
	}, nil
}

func (s *stringType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if s.reflectValue == nil {
		return s.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return s.reflectValue.CanConvert(typ.ReflectType())
}

func (s *stringType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if s.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !s.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := s.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (s *stringType) StringValue() (string, error) {
	if s.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return s.reflectValue.Interface().(string), nil
}

func (s *stringType) SetStringValue(val string) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	s.reflectValue.Set(reflect.ValueOf(val))
	return nil
}
