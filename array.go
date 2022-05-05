package reflector

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Array interface {
	Type
	Elem() Type
	Len() int
	Get(index int) (any, error)
	Set(index int, val any) error
	Slice(low, high int) (any, error)
	Copy(dst any) (int, error)
}

type arrayType struct {
	elem Type

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (a *arrayType) Name() string {
	var builder strings.Builder
	builder.WriteString("[")
	builder.WriteString(fmt.Sprintf("%d", a.Len()))
	builder.WriteString("]")
	builder.WriteString(a.elem.Name())
	return builder.String()
}

func (a *arrayType) PackageName() string {
	return ""
}

func (a *arrayType) PackagePath() string {
	return ""
}

func (a *arrayType) CanSet() bool {
	if a.reflectValue == nil {
		return false
	}

	return a.reflectValue.CanSet()
}

func (a *arrayType) HasValue() bool {
	return a.reflectValue != nil
}

func (a *arrayType) Value() (any, error) {
	if a.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return a.reflectValue.Interface(), nil
}

func (a *arrayType) SetValue(val any) error {
	if !a.CanSet() {
		return errors.New("value cannot be set")
	}

	a.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (a *arrayType) Parent() Type {
	return a.parent
}

func (a *arrayType) ReflectType() reflect.Type {
	return a.reflectType
}

func (a *arrayType) ReflectValue() *reflect.Value {
	return a.reflectValue
}

func (a *arrayType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return a.reflectType == another.ReflectType()
}

func (a *arrayType) IsInstantiable() bool {
	return true
}

func (a *arrayType) Instantiate() (Value, error) {
	return &value{
		reflect.New(a.reflectType),
	}, nil
}

func (a *arrayType) Elem() Type {
	return a.elem
}

func (a *arrayType) Len() int {
	return a.reflectType.Len()
}

func (a *arrayType) Get(index int) (any, error) {
	if a.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if index < 0 || index >= a.Len() {
		return nil, errors.New("array index out of range")
	}

	return a.reflectValue.Index(index).Interface(), nil
}

func (a *arrayType) Set(index int, val any) error {
	if !a.CanSet() {
		return errors.New("value cannot be set")
	}

	if index < 0 || index >= a.Len() {
		return errors.New("array index out of range")
	}

	a.reflectValue.Index(index).Set(reflect.ValueOf(val))
	return nil
}

func (a *arrayType) Slice(low, high int) (any, error) {
	if a.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if low < 0 || low > a.Len() {
		return nil, errors.New("array index out of range")
	}

	if high < 0 || high > a.Len() {
		return nil, errors.New("array index out of range")
	}

	return a.reflectValue.Slice(low, high).Interface(), nil
}

func (a *arrayType) Copy(dst any) (int, error) {
	if a.reflectValue == nil {
		return -1, errors.New("value reference is nil")
	}

	// TODO BUG: It causes app to crash
	return reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(a.reflectValue.Slice(0, a.Len()))), nil
}
