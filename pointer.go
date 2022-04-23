package reflector

import "reflect"

type Pointer interface {
	Type
	Elem() Type
}

type pointer struct {
	base Type

	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (p *pointer) HasReference() bool {
	return p.reflectValue != nil
}

func (p *pointer) ReflectType() reflect.Type {
	return p.reflectType
}

func (p *pointer) ReflectValue() *reflect.Value {
	return p.reflectValue
}

func (p *pointer) Elem() Type {
	return p.base
}
