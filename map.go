package reflector

import (
	"reflect"
)

type Entry interface {
	Key() any
	Value() any
}

type entry struct {
	key any
	val any
}

func (e entry) Key() any {
	return e.key
}

func (e entry) Value() any {
	return e.val
}

type Map interface {
	Type
	Instantiable
	Key() Type
	Value() Type
	Len() int
	Get(key any) (any, bool)
	Contains(key any) bool
	Put(key any, val any)
	KeySet() []any
	ValueSet() []any
	EntrySet() []Entry
}

type mapType struct {
	key  Type
	elem Type

	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (m *mapType) Name() string {
	return m.reflectType.Name()
}

func (m *mapType) PackageName() string {
	return ""
}

func (m *mapType) HasReference() bool {
	return m.reflectValue != nil
}

func (m *mapType) ReflectType() reflect.Type {
	return m.reflectType
}

func (m *mapType) ReflectValue() *reflect.Value {
	return m.reflectValue
}

func (m *mapType) Key() Type {
	return m.key
}

func (m *mapType) Value() Type {
	return m.elem
}

func (m *mapType) Len() int {
	if m.reflectValue == nil {
		return -1
	}

	return m.reflectValue.Len()
}

func (m *mapType) Contains(key any) bool {
	_, ok := m.Get(key)
	return ok
}

func (m *mapType) KeySet() []any {
	if m.reflectValue == nil {
		return nil
	}

	keySet := make([]any, 0)
	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		keySet = append(keySet, key.Interface())
	}

	return keySet
}

func (m *mapType) ValueSet() []any {
	if m.reflectValue == nil {
		return nil
	}

	valueSet := make([]any, 0)
	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		value := m.reflectValue.MapIndex(key)
		valueSet = append(valueSet, value.Interface())
	}

	return valueSet
}

func (m *mapType) EntrySet() []Entry {
	if m.reflectValue == nil {
		return nil
	}

	valueSet := make([]Entry, 0)
	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		value := m.reflectValue.MapIndex(key)
		valueSet = append(valueSet, entry{key: key.Interface(), val: value.Interface()})
	}

	return valueSet
}

func (m *mapType) Get(key any) (any, bool) {
	if m.reflectValue == nil {
		return nil, false
	}

	val := m.reflectValue.MapIndex(reflect.ValueOf(key))

	if val.Kind() == reflect.Invalid {
		return nil, false
	}

	return val.Interface(), true
}

func (m *mapType) Put(key any, val any) {
	if m.reflectValue == nil {
		return
	}

	m.reflectValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
}

func (m *mapType) Instantiate() Value {
	ptr := reflect.New(m.reflectType).Interface()
	emptyMap := reflect.MakeMapWithSize(reflect.MapOf(m.key.ReflectType(), m.elem.ReflectType()), 0)
	reflect.ValueOf(ptr).Elem().Set(emptyMap)
	return &value{
		reflect.ValueOf(ptr),
	}
}
