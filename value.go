package reflector

import "reflect"

type Value interface {
	Val() any
	Elem() any
}

type value struct {
	reflectValue reflect.Value
}

func (v *value) Val() any {
	return v.reflectValue.Interface()
}

func (v *value) Elem() any {
	return v.reflectValue.Elem().Interface()
}
