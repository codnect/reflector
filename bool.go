package reflector

import "reflect"

type Boolean interface {
	CanSet() bool
	Value() bool
	SetValue(val bool)
}

type booleanType struct {
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (b *booleanType) Name() string {
	return b.reflectType.Name()
}

func (b *booleanType) PackageName() string {
	return ""
}

func (b *booleanType) HasReference() bool {
	return b.reflectValue != nil
}

func (b *booleanType) ReflectType() reflect.Type {
	return b.reflectType
}

func (b *booleanType) ReflectValue() *reflect.Value {
	return b.reflectValue
}

func (b *booleanType) CanSet() bool {
	if b.reflectValue == nil {
		return false
	}

	return b.reflectValue.CanSet()
}

func (b *booleanType) Value() bool {
	if b.reflectValue == nil {
		return false
	}

	return b.reflectValue.Interface().(bool)
}

func (b *booleanType) SetValue(val bool) {
	if b.reflectValue == nil {
		return
	}

	b.reflectValue.Set(reflect.ValueOf(val))
}

func (b *booleanType) Instantiate() any {
	return reflect.New(b.reflectType).Interface()
}
