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
	assert.True(t, booleanType.IsInstantiable())
	assert.True(t, booleanType.Compare(TypeOf[bool]()))
	assert.False(t, booleanType.Compare(TypeOf[*bool]()))
	assert.False(t, booleanType.Compare(TypeOf[string]()))

	value, err := booleanType.Value()
	assert.False(t, value.(bool))
	assert.NotNil(t, err)

	err = booleanType.SetValue(true)
	assert.NotNil(t, err)

	booleanValue, err := booleanType.BooleanValue()
	assert.False(t, booleanValue)
	assert.NotNil(t, err)

	err = booleanType.SetBooleanValue(false)
	assert.NotNil(t, err)

	newBool, err := booleanType.Instantiate()
	assert.NotNil(t, newBool)
	assert.Nil(t, err)

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

	assert.False(t, ptrType.CanSet())
	assert.True(t, ptrType.IsInstantiable())
	assert.False(t, ptrType.Compare(TypeOf[bool]()))
	assert.True(t, ptrType.Compare(TypeOf[*bool]()))
	assert.False(t, ptrType.Compare(TypeOf[string]()))

	value, err := ptrType.Value()
	assert.Nil(t, value)
	assert.NotNil(t, err)

	err = ptrType.SetValue(nil)
	assert.NotNil(t, err)

	value, err = ptrType.Value()
	assert.Nil(t, value)
	assert.NotNil(t, err)

	newBool, err := ptrType.Instantiate()
	assert.NotNil(t, newBool)
	assert.Nil(t, err)

	boolPtrVal, ok := newBool.Val().(*bool)
	assert.True(t, ok)
	assert.False(t, *boolPtrVal)

	boolVal, ok := newBool.Elem().(bool)
	assert.True(t, ok)
	assert.False(t, boolVal)

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
	assert.True(t, booleanType.IsInstantiable())
	assert.True(t, booleanType.Compare(TypeOf[bool]()))
	assert.False(t, booleanType.Compare(TypeOf[*bool]()))
	assert.False(t, booleanType.Compare(TypeOf[string]()))

	value, err = booleanType.Value()
	assert.False(t, value.(bool))
	assert.NotNil(t, err)

	err = booleanType.SetValue(true)
	assert.NotNil(t, err)

	booleanValue, err := booleanType.BooleanValue()
	assert.False(t, booleanValue)
	assert.NotNil(t, err)

	err = booleanType.SetBooleanValue(false)
	assert.NotNil(t, err)

	newBool, err = booleanType.Instantiate()
	assert.NotNil(t, newBool)
	assert.Nil(t, err)

	boolPtrVal, ok = newBool.Val().(*bool)
	assert.True(t, ok)
	assert.False(t, *boolPtrVal)

	boolVal, ok = newBool.Elem().(bool)
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
	assert.True(t, booleanType.IsInstantiable())
	assert.True(t, booleanType.Compare(TypeOf[bool]()))
	assert.False(t, booleanType.Compare(TypeOf[*bool]()))
	assert.False(t, booleanType.Compare(TypeOf[string]()))

	value, err := booleanType.Value()
	assert.True(t, value.(bool))
	assert.Nil(t, err)

	err = booleanType.SetValue(true)
	assert.NotNil(t, err)

	booleanValue, err := booleanType.BooleanValue()
	assert.True(t, booleanValue)
	assert.Nil(t, err)

	err = booleanType.SetBooleanValue(false)
	assert.NotNil(t, err)

	newBool, err := booleanType.Instantiate()
	assert.NotNil(t, newBool)
	assert.Nil(t, err)

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

	assert.False(t, ptrType.CanSet())
	assert.True(t, ptrType.IsInstantiable())
	assert.False(t, ptrType.Compare(TypeOf[bool]()))
	assert.True(t, ptrType.Compare(TypeOf[*bool]()))
	assert.False(t, ptrType.Compare(TypeOf[string]()))

	value, err := ptrType.Value()
	assert.NotNil(t, value)
	assert.Nil(t, err)

	err = ptrType.SetValue(&val)
	assert.NotNil(t, err)

	value, err = ptrType.Value()
	assert.NotNil(t, value)
	assert.Nil(t, err)

	newBool, err := ptrType.Instantiate()
	assert.NotNil(t, newBool)
	assert.Nil(t, err)

	boolPtrVal, ok := newBool.Val().(*bool)
	assert.True(t, ok)
	assert.False(t, *boolPtrVal)

	boolVal, ok := newBool.Elem().(bool)
	assert.True(t, ok)
	assert.False(t, boolVal)

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
	assert.True(t, booleanType.IsInstantiable())
	assert.True(t, booleanType.Compare(TypeOf[bool]()))
	assert.False(t, booleanType.Compare(TypeOf[*bool]()))
	assert.False(t, booleanType.Compare(TypeOf[string]()))

	value, err = booleanType.Value()
	assert.True(t, value.(bool))
	assert.Nil(t, err)

	err = booleanType.SetValue(false)
	assert.Nil(t, err)

	booleanValue, err := booleanType.BooleanValue()
	assert.False(t, booleanValue)
	assert.Nil(t, err)

	err = booleanType.SetBooleanValue(false)
	assert.Nil(t, err)

	newBool, err = booleanType.Instantiate()
	assert.NotNil(t, newBool)
	assert.Nil(t, err)

	boolPtrVal, ok = newBool.Val().(*bool)
	assert.True(t, ok)
	assert.False(t, *boolPtrVal)

	boolVal, ok = newBool.Elem().(bool)
	assert.True(t, ok)
	assert.False(t, boolVal)
}
