package reflector

import "reflect"

type String interface {
	CanSet() bool
	Value() string
	SetValue(val string)
}

type stringType struct {
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *stringType) Name() string {
	return s.reflectType.Name()
}

func (s *stringType) PackageName() string {
	return ""
}

func (s *stringType) HasReference() bool {
	return s.reflectValue != nil
}

func (s *stringType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *stringType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *stringType) CanSet() bool {
	if s.reflectValue == nil {
		return false
	}

	return s.reflectValue.CanSet()
}

func (s *stringType) Value() string {
	if s.reflectValue == nil {
		return ""
	}

	return s.reflectValue.Interface().(string)
}

func (s *stringType) SetValue(val string) {
	if s.reflectValue == nil {
		return
	}

	s.reflectValue.Set(reflect.ValueOf(val))
}

func (s *stringType) Instantiate() any {
	return reflect.New(s.reflectType).Interface()
}
