package reflector

import "reflect"

type Array interface {
	Type
	Instantiable
	Value() any
	Index(i int) any
	Set(i int, v any)
	Len() int
	Elem() Type
}

type arrayType struct {
	elem Type

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (a *arrayType) Name() string {
	return a.reflectType.Name()
}

func (a *arrayType) PackageName() string {
	return ""
}

func (a *arrayType) HasValue() bool {
	return a.reflectValue != nil
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

func (a *arrayType) Value() any {
	return a.reflectValue.Interface()
}

func (a *arrayType) Set(i int, v any) {

}

func (a *arrayType) Index(i int) any {
	return a.reflectValue.Index(i).Interface()
}

func (a *arrayType) Len() int {
	return a.reflectType.Len()
}

func (a *arrayType) Elem() Type {
	return a.elem
}

func (a *arrayType) Instantiate() Value {
	return &value{
		reflect.New(a.reflectType),
	}
}
