package reflector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeOfBoolean(t *testing.T) {
	typ := TypeOf[bool]()
	assert.True(t, IsBoolean(typ))
	assert.Equal(t, "bool", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	booleanType, isBoolean := ToBoolean(typ)

	assert.NotNil(t, booleanType)
	assert.True(t, isBoolean)

	assert.False(t, booleanType.CanSet())

	value, err := booleanType.Value()
	assert.False(t, value)
	assert.NotNil(t, err)

	err = booleanType.SetValue(true)
	assert.NotNil(t, err)

	newBool := booleanType.Instantiate()
	assert.NotNil(t, newBool)

	boolPtrVal, ok := newBool.Val().(*bool)
	assert.True(t, ok)
	assert.False(t, *boolPtrVal)

	boolVal, ok := newBool.Elem().(bool)
	assert.True(t, ok)
	assert.False(t, boolVal)
}

func TestTypeOfBooleanPointer(t *testing.T) {
	ptrType := TypeOf[*bool]()
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*bool", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBoolean(typ))
	assert.Equal(t, "bool", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	booleanType, isBoolean := ToBoolean(typ)

	assert.NotNil(t, booleanType)
	assert.True(t, isBoolean)

	assert.False(t, booleanType.CanSet())

	value, err := booleanType.Value()
	assert.False(t, value)
	assert.NotNil(t, err)

	err = booleanType.SetValue(true)
	assert.NotNil(t, err)

	newBool := booleanType.Instantiate()
	assert.NotNil(t, newBool)

	boolPtrVal, ok := newBool.Val().(*bool)
	assert.True(t, ok)
	assert.False(t, *boolPtrVal)

	boolVal, ok := newBool.Elem().(bool)
	assert.True(t, ok)
	assert.False(t, boolVal)
}

func TestTypeOfBooleanObject(t *testing.T) {
	val := true

	typ := TypeOfAny(val)
	assert.True(t, IsBoolean(typ))
	assert.Equal(t, "bool", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	booleanType, isBoolean := ToBoolean(typ)

	assert.NotNil(t, booleanType)
	assert.True(t, isBoolean)

	assert.False(t, booleanType.CanSet())

	value, err := booleanType.Value()
	assert.True(t, value)
	assert.Nil(t, err)

	err = booleanType.SetValue(false)
	assert.NotNil(t, err)
	assert.True(t, val)

	newBool := booleanType.Instantiate()
	assert.NotNil(t, newBool)

	boolPtrVal, ok := newBool.Val().(*bool)
	assert.True(t, ok)
	assert.False(t, *boolPtrVal)

	boolVal, ok := newBool.Elem().(bool)
	assert.True(t, ok)
	assert.False(t, boolVal)
}

func TestTypeOfBooleanObjectPointer(t *testing.T) {
	val := true

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*bool", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBoolean(typ))
	assert.Equal(t, "bool", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	booleanType, isBoolean := ToBoolean(typ)

	assert.NotNil(t, booleanType)
	assert.True(t, isBoolean)

	assert.True(t, booleanType.CanSet())

	value, err := booleanType.Value()
	assert.True(t, value)
	assert.Nil(t, err)

	err = booleanType.SetValue(false)
	assert.Nil(t, err)
	assert.False(t, val)

	newBool := booleanType.Instantiate()
	assert.NotNil(t, newBool)

	boolPtrVal, ok := newBool.Val().(*bool)
	assert.True(t, ok)
	assert.False(t, *boolPtrVal)

	boolVal, ok := newBool.Elem().(bool)
	assert.True(t, ok)
	assert.False(t, boolVal)
}
