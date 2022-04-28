package reflector

import (
	"reflect"
)

type ChanDirection int

const (
	SEND ChanDirection = 1 << iota
	RECEIVE
)

type Chan interface {
	Type
	Direction() ChanDirection
}

type chanType struct {
	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (i *chanType) Name() string {
	return i.reflectType.Name()
}

func (i *chanType) PackageName() string {
	return ""
}

func (i *chanType) HasValue() bool {
	return i.reflectValue != nil
}

func (i *chanType) Parent() Type {
	return i.parent
}

func (i *chanType) ReflectType() reflect.Type {
	return i.reflectType
}

func (i *chanType) ReflectValue() *reflect.Value {
	return i.reflectValue
}

func (i *chanType) Direction() ChanDirection {
	return SEND
}

func (i *chanType) Elem() Type {
	return nil
}
