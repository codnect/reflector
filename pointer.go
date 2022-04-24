package reflector

import (
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

func (p *pointer) HasValue() bool {
	return p.reflectValue != nil
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
