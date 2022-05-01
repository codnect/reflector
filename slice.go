package reflector

import (
	"errors"
	"reflect"
	"strings"
)

type Slice interface {
	Type
	Instantiable
	CanSet() bool
	Elem() Type
	Value() (any, error)
	SetValue(val any) error
	Len() (int, error)
	Cap() (int, error)
	Get(index int) (any, error)
	Set(index int, val any) error
	Append(values ...any)
	Slice(low, high int) (any, error)
	Copy(dst any) (int, error)
}

type sliceType struct {
	elem Type

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *sliceType) Name() string {
	var builder strings.Builder
	builder.WriteString("[]")
	builder.WriteString(s.elem.Name())
	return builder.String()
}

func (s *sliceType) PackageName() string {
	return ""
}

func (s *sliceType) PackagePath() string {
	return ""
}

func (s *sliceType) CanSet() bool {
	if s.reflectValue == nil {
		return false
	}

	return s.reflectValue.CanSet()
}

func (s *sliceType) Elem() Type {
	return s.elem
}

func (s *sliceType) Value() (any, error) {
	if s.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return s.reflectValue.Interface(), nil
}

func (s *sliceType) SetValue(val any) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	s.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (s *sliceType) Len() (int, error) {
	if s.reflectValue == nil {
		return -1, errors.New("value reference is nil")
	}

	return s.reflectValue.Len(), nil
}

func (s *sliceType) Cap() (int, error) {
	if s.reflectValue == nil {
		return -1, errors.New("value reference is nil")
	}

	return s.reflectValue.Cap(), nil
}

func (s *sliceType) Get(index int) (any, error) {
	if s.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if index < 0 || index >= s.reflectValue.Len() {
		return nil, errors.New("array index out of range")
	}

	return s.reflectValue.Index(index).Interface(), nil
}

func (s *sliceType) Set(index int, val any) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	if index < 0 || index >= s.reflectValue.Len() {
		return errors.New("array index out of range")
	}

	s.reflectValue.Index(index).Set(reflect.ValueOf(val))
	return nil
}

func (s *sliceType) Append(values ...any) {
	if !s.CanSet() {
		return
	}

	for _, value := range values {
		*s.reflectValue = reflect.Append(*s.reflectValue, reflect.ValueOf(value))
	}
}

func (s *sliceType) Slice(low, high int) (any, error) {
	if s.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if low < 0 || low > s.reflectValue.Len() {
		return nil, errors.New("array index out of range")
	}

	if high < 0 || high > s.reflectValue.Len() {
		return nil, errors.New("array index out of range")
	}

	return s.reflectValue.Slice(low, high).Interface(), nil
}

func (s *sliceType) Copy(dst any) (int, error) {
	if s.reflectValue == nil {
		return -1, errors.New("value reference is nil")
	}

	return reflect.Copy(reflect.ValueOf(dst), *s.reflectValue), nil
}

func (s *sliceType) HasValue() bool {
	return s.reflectValue != nil
}

func (s *sliceType) Parent() Type {
	return s.parent
}

func (s *sliceType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *sliceType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *sliceType) Instantiate() Value {
	return &value{
		reflect.New(s.reflectType),
	}
}
