package reflector

import "testing"
import "github.com/stretchr/testify/assert"

func TestTypeOfBoolean(t *testing.T) {
	typ := TypeOf[bool]()
	assert.True(t, IsBoolean(typ))
	assert.Equal(t, "bool", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasReference())
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
	pointerBoolType := TypeOf[*bool]()
	assert.True(t, IsPointer(pointerBoolType))
	ptr, isPtr := ToPointer(pointerBoolType)
	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	/*
		assert.True(t, IsBoolean(typ))
		assert.Equal(t, "bool", typ.Name())
		assert.Equal(t, "", typ.PackageName())

		assert.False(t, typ.HasReference())
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

	*/
}
