package reflector

import (
	"errors"
	"reflect"
	"strings"
)

type Pointer interface {
	Type
	Elem() Type
}

type pointer struct {
	base Type

	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (p *pointer) Name() string {
	var builder strings.Builder
	builder.WriteString("*")
	builder.WriteString(p.base.Name())
	return builder.String()
}

func (p *pointer) PackageName() string {
	return p.base.PackageName()
}

func (p *pointer) PackagePath() string {
	return p.base.PackagePath()
}

func (p *pointer) HasValue() bool {
	return p.reflectValue != nil
}

func (p *pointer) CanSet() bool {
	if p.reflectValue == nil {
		return false
	}

	return p.reflectValue.CanSet()
}

func (p *pointer) Value() (any, error) {
	if p.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	return p.reflectValue.Interface(), nil
}

func (p *pointer) SetValue(val any) error {
	if !p.CanSet() {
		return errors.New("value cannot be set")
	}

	p.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (p *pointer) ReflectType() reflect.Type {
	return p.reflectType
}

func (p *pointer) ReflectValue() *reflect.Value {
	return p.reflectValue
}

func (p *pointer) Parent() Type {
	return nil
}

func (p *pointer) Elem() Type {
	return p.base
}

func (p *pointer) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return p.reflectType == another.ReflectType()
}

func (p *pointer) IsInstantiable() bool {
	return p.base.IsInstantiable()
}

func (p *pointer) Instantiate() (Value, error) {
	return p.base.Instantiate()
}

func (p *pointer) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if p.reflectValue == nil {
		return p.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return p.reflectValue.CanConvert(typ.ReflectType())
}

func (p *pointer) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if p.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !p.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := p.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}
