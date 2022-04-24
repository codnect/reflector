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
	CanSet() bool
	Key() Type
	Value() Type
	Len() (int, error)
	Get(key any) (any, error)
	Put(key any, val any) error
	KeySet() ([]any, error)
	ValueSet() ([]any, error)
	EntrySet() []Entry
}

type mapType struct {
	key  Type
	elem Type

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (m *mapType) Name() string {
	return m.reflectType.Name()
}

func (m *mapType) PackageName() string {
	return ""
}

func (m *mapType) HasValue() bool {
	return m.reflectValue != nil
}

func (m *mapType) Parent() Type {
	return m.parent
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

func (m *mapType) CanSet() bool {
	return true
}

func (m *mapType) Len() (int, error) {
	if m.reflectValue == nil {
		return -1, nil
	}

	return m.reflectValue.Len(), nil
}

func (m *mapType) KeySet() ([]any, error) {
	if m.reflectValue == nil {
		return nil, nil
	}

	keySet := make([]any, 0)
	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		keySet = append(keySet, key.Interface())
	}

	return keySet, nil
}

func (m *mapType) ValueSet() ([]any, error) {
	if m.reflectValue == nil {
		return nil, nil
	}

	valueSet := make([]any, 0)
	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		value := m.reflectValue.MapIndex(key)
		valueSet = append(valueSet, value.Interface())
	}

	return valueSet, nil
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

func (m *mapType) Get(key any) (any, error) {
	if m.reflectValue == nil {
		return nil, nil
	}

	val := m.reflectValue.MapIndex(reflect.ValueOf(key))

	if val.Kind() == reflect.Invalid {
		return nil, nil
	}

	return val.Interface(), nil
}

func (m *mapType) Put(key any, val any) error {
	if m.reflectValue == nil {
		return nil
	}

	m.reflectValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	return nil
}

func (m *mapType) Instantiate() Value {
	ptr := reflect.New(m.reflectType).Interface()
	emptyMap := reflect.MakeMapWithSize(reflect.MapOf(m.key.ReflectType(), m.elem.ReflectType()), 0)
	reflect.ValueOf(ptr).Elem().Set(emptyMap)
	return &value{
		reflect.ValueOf(ptr),
	}
}
