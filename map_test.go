package reflector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeOfMap(t *testing.T) {
	typ := TypeOf[map[string]bool]()
	assert.True(t, IsMap(typ))
	assert.Equal(t, "map[string]bool", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	mapType, isMap := ToMap(typ)

	assert.NotNil(t, mapType)
	assert.True(t, isMap)

	assert.False(t, mapType.CanSet())

	value, err := mapType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = mapType.SetValue(map[string]any{"key2": true})
	assert.NotNil(t, err)

	len, err := mapType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	v, err := mapType.Get("key1")
	assert.Nil(t, v)
	assert.NotNil(t, err)

	exists, err := mapType.Contains("key1")
	assert.NotNil(t, err)
	assert.False(t, exists)

	keySet, err := mapType.KeySet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	valueSet, err := mapType.ValueSet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	entrySet, err := mapType.EntrySet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	err = mapType.Put("key2", true)
	assert.NotNil(t, err)

	exists, err = mapType.Contains("key2")
	assert.NotNil(t, err)
	assert.False(t, exists)

	value, err = mapType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	len, err = mapType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	keySet, err = mapType.KeySet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	valueSet, err = mapType.ValueSet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	entrySet, err = mapType.EntrySet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	err = mapType.Delete("key1")
	assert.NotNil(t, err)

	value, err = mapType.Value()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	len, err = mapType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	keySet, err = mapType.KeySet()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	valueSet, err = mapType.ValueSet()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	entrySet, err = mapType.EntrySet()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	err = mapType.Clear()
	assert.NotNil(t, err)

	value, err = mapType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	len, err = mapType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	keySet, err = mapType.KeySet()
	assert.Len(t, keySet, 0)

	valueSet, err = mapType.ValueSet()
	assert.Len(t, valueSet, 0)

	entrySet, err = mapType.EntrySet()
	assert.Len(t, entrySet, 0)

	newMap, _ := mapType.Instantiate()

	mapPtrVal, ok := newMap.Val().(*map[string]bool)
	assert.True(t, ok)
	assert.Empty(t, *mapPtrVal)

	mapVal, ok := newMap.Elem().(map[string]bool)
	assert.True(t, ok)
	assert.Empty(t, mapVal)
}

func TestTypeOfMapPointer(t *testing.T) {
	ptrType := TypeOf[*map[string]bool]()
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*map[string]bool", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsMap(typ))
	assert.Equal(t, "map[string]bool", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	mapType, isMap := ToMap(typ)

	assert.NotNil(t, mapType)
	assert.True(t, isMap)

	assert.False(t, mapType.CanSet())

	value, err := mapType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = mapType.SetValue(map[string]any{"key2": true})
	assert.NotNil(t, err)

	len, err := mapType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	v, err := mapType.Get("key1")
	assert.Nil(t, v)
	assert.NotNil(t, err)

	exists, err := mapType.Contains("key1")
	assert.NotNil(t, err)
	assert.False(t, exists)

	keySet, err := mapType.KeySet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	valueSet, err := mapType.ValueSet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	entrySet, err := mapType.EntrySet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	err = mapType.Put("key2", true)
	assert.NotNil(t, err)

	exists, err = mapType.Contains("key2")
	assert.NotNil(t, err)
	assert.False(t, exists)

	value, err = mapType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	len, err = mapType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	keySet, err = mapType.KeySet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	valueSet, err = mapType.ValueSet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	entrySet, err = mapType.EntrySet()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	err = mapType.Delete("key1")
	assert.NotNil(t, err)

	value, err = mapType.Value()
	assert.Nil(t, v)
	assert.NotNil(t, err)

	len, err = mapType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	keySet, err = mapType.KeySet()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	valueSet, err = mapType.ValueSet()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	entrySet, err = mapType.EntrySet()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	err = mapType.Clear()
	assert.NotNil(t, err)

	value, err = mapType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	len, err = mapType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	keySet, err = mapType.KeySet()
	assert.Len(t, keySet, 0)

	valueSet, err = mapType.ValueSet()
	assert.Len(t, valueSet, 0)

	entrySet, err = mapType.EntrySet()
	assert.Len(t, entrySet, 0)

	newMap, _ := mapType.Instantiate()

	mapPtrVal, ok := newMap.Val().(*map[string]bool)
	assert.True(t, ok)
	assert.Empty(t, *mapPtrVal)

	mapVal, ok := newMap.Elem().(map[string]bool)
	assert.True(t, ok)
	assert.Empty(t, mapVal)
}

func TestTypeOfMapObject(t *testing.T) {
	val := map[string]any{"key1": "value1"}

	typ := TypeOfAny(val)
	assert.True(t, IsMap(typ))
	assert.Equal(t, "map[string]any", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	mapType, isMap := ToMap(typ)

	assert.NotNil(t, mapType)
	assert.True(t, isMap)

	assert.False(t, mapType.CanSet())

	value, err := mapType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)

	err = mapType.SetValue(map[string]any{"key2": true})
	assert.NotNil(t, err)
	assert.NotEqual(t, map[string]any{"key2": true}, val)

	assert.Len(t, value, 1)
	len, err := mapType.Len()
	assert.Equal(t, 1, len)
	assert.Nil(t, err)

	v, err := mapType.Get("key1")
	assert.Equal(t, "value1", v.(string))
	assert.Nil(t, err)

	exists, err := mapType.Contains("key1")
	assert.Nil(t, err)
	assert.True(t, exists)

	keySet, err := mapType.KeySet()
	assert.Len(t, keySet, 1)
	assert.Contains(t, keySet, "key1")

	valueSet, err := mapType.ValueSet()
	assert.Len(t, valueSet, 1)
	assert.Contains(t, valueSet, "value1")

	entrySet, err := mapType.EntrySet()
	assert.Len(t, entrySet, 1)
	assert.Equal(t, "key1", entrySet[0].Key())
	assert.Equal(t, "value1", entrySet[0].Value())

	err = mapType.Put("key2", true)
	assert.Nil(t, err)

	exists, err = mapType.Contains("key2")
	assert.Nil(t, err)
	assert.True(t, exists)

	value, err = mapType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)

	assert.Len(t, value, 2)
	len, err = mapType.Len()
	assert.Equal(t, 2, len)
	assert.Nil(t, err)

	keySet, err = mapType.KeySet()
	assert.Len(t, keySet, 2)
	assert.Contains(t, keySet, "key1")
	assert.Contains(t, keySet, "key2")

	valueSet, err = mapType.ValueSet()
	assert.Len(t, valueSet, 2)
	assert.Contains(t, valueSet, "value1")
	assert.Contains(t, valueSet, true)

	entrySet, err = mapType.EntrySet()
	assert.Len(t, entrySet, 2)

	err = mapType.Delete("key1")
	assert.Nil(t, err)

	exists, err = mapType.Contains("key2")
	assert.Nil(t, err)
	assert.True(t, exists)

	value, err = mapType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)

	assert.Len(t, value, 1)
	len, err = mapType.Len()
	assert.Equal(t, 1, len)
	assert.Nil(t, err)

	keySet, err = mapType.KeySet()
	assert.Len(t, keySet, 1)
	assert.Contains(t, keySet, "key2")

	valueSet, err = mapType.ValueSet()
	assert.Len(t, valueSet, 1)
	assert.Contains(t, valueSet, true)

	entrySet, err = mapType.EntrySet()
	assert.Len(t, entrySet, 1)
	assert.Equal(t, "key2", entrySet[0].Key())
	assert.Equal(t, true, entrySet[0].Value())

	err = mapType.Clear()
	assert.Nil(t, err)

	exists, err = mapType.Contains("key2")
	assert.NotNil(t, v)
	assert.False(t, exists)

	value, err = mapType.Value()
	assert.Empty(t, value)
	assert.Nil(t, err)

	assert.Len(t, value, 0)
	len, err = mapType.Len()
	assert.Equal(t, 0, len)
	assert.Nil(t, err)

	keySet, err = mapType.KeySet()
	assert.Len(t, keySet, 0)

	valueSet, err = mapType.ValueSet()
	assert.Len(t, valueSet, 0)

	entrySet, err = mapType.EntrySet()
	assert.Len(t, entrySet, 0)

	newMap, _ := mapType.Instantiate()

	mapPtrVal, ok := newMap.Val().(*map[string]any)
	assert.True(t, ok)
	assert.Empty(t, *mapPtrVal)

	mapVal, ok := newMap.Elem().(map[string]any)
	assert.True(t, ok)
	assert.Empty(t, mapVal)
}

func TestTypeOfMapObjectPointer(t *testing.T) {
	val := map[string]any{"key1": "value1"}

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*map[string]any", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsMap(typ))
	assert.Equal(t, "map[string]any", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	mapType, isMap := ToMap(typ)

	assert.NotNil(t, mapType)
	assert.True(t, isMap)

	assert.True(t, mapType.CanSet())

	value, err := mapType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)

	err = mapType.SetValue(map[string]any{"key2": true})
	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"key2": true}, val)

	val["key1"] = "value1"
	delete(val, "key2")

	assert.Len(t, value, 1)
	len, err := mapType.Len()
	assert.Equal(t, 1, len)
	assert.Nil(t, err)

	v, err := mapType.Get("key1")
	assert.Equal(t, "value1", v.(string))
	assert.Nil(t, err)

	exists, err := mapType.Contains("key1")
	assert.Nil(t, err)
	assert.True(t, exists)

	keySet, err := mapType.KeySet()
	assert.Len(t, keySet, 1)
	assert.Contains(t, keySet, "key1")

	valueSet, err := mapType.ValueSet()
	assert.Len(t, valueSet, 1)
	assert.Contains(t, valueSet, "value1")

	entrySet, err := mapType.EntrySet()
	assert.Len(t, entrySet, 1)
	assert.Equal(t, "key1", entrySet[0].Key())
	assert.Equal(t, "value1", entrySet[0].Value())

	err = mapType.Put("key2", true)
	assert.Nil(t, err)

	exists, err = mapType.Contains("key2")
	assert.Nil(t, err)
	assert.True(t, exists)

	value, err = mapType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)

	assert.Len(t, value, 2)
	len, err = mapType.Len()
	assert.Equal(t, 2, len)
	assert.Nil(t, err)

	keySet, err = mapType.KeySet()
	assert.Len(t, keySet, 2)
	assert.Contains(t, keySet, "key1")
	assert.Contains(t, keySet, "key2")

	valueSet, err = mapType.ValueSet()
	assert.Len(t, valueSet, 2)
	assert.Contains(t, valueSet, "value1")
	assert.Contains(t, valueSet, true)

	entrySet, err = mapType.EntrySet()
	assert.Len(t, entrySet, 2)

	err = mapType.Delete("key1")
	assert.Nil(t, err)

	value, err = mapType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)

	assert.Len(t, value, 1)
	len, err = mapType.Len()
	assert.Equal(t, 1, len)
	assert.Nil(t, err)

	keySet, err = mapType.KeySet()
	assert.Len(t, keySet, 1)
	assert.Contains(t, keySet, "key2")

	valueSet, err = mapType.ValueSet()
	assert.Len(t, valueSet, 1)
	assert.Contains(t, valueSet, true)

	entrySet, err = mapType.EntrySet()
	assert.Len(t, entrySet, 1)
	assert.Equal(t, "key2", entrySet[0].Key())
	assert.Equal(t, true, entrySet[0].Value())

	err = mapType.Clear()
	assert.Nil(t, err)

	value, err = mapType.Value()
	assert.Empty(t, value)
	assert.Nil(t, err)

	assert.Len(t, value, 0)
	len, err = mapType.Len()
	assert.Equal(t, 0, len)
	assert.Nil(t, err)

	keySet, err = mapType.KeySet()
	assert.Len(t, keySet, 0)

	valueSet, err = mapType.ValueSet()
	assert.Len(t, valueSet, 0)

	entrySet, err = mapType.EntrySet()
	assert.Len(t, entrySet, 0)

	newMap, _ := mapType.Instantiate()

	mapPtrVal, ok := newMap.Val().(*map[string]any)
	assert.True(t, ok)
	assert.Empty(t, *mapPtrVal)

	mapVal, ok := newMap.Elem().(map[string]any)
	assert.True(t, ok)
	assert.Empty(t, mapVal)
}
