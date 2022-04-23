package reflector

import "reflect"

type Struct interface {
	Type
	Instantiable
}

type structType struct {
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *structType) HasReference() bool {
	return s.reflectValue != nil
}

func (s *structType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *structType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *structType) Instantiate() any {
	return reflect.New(s.reflectType).Interface()
}
