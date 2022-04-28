package reflector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeOfString(t *testing.T) {
	typ := TypeOf[string]()
	assert.True(t, IsString(typ))
	assert.Equal(t, "string", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	stringType, isString := ToString(typ)

	assert.NotNil(t, stringType)
	assert.True(t, isString)

	assert.False(t, stringType.CanSet())

	value, err := stringType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = stringType.SetValue("hello")
	assert.NotNil(t, err)

	newString := stringType.Instantiate()
	assert.NotNil(t, newString)

	stringPtrVal, ok := newString.Val().(*string)
	assert.True(t, ok)
	assert.Empty(t, *stringPtrVal)

	stringVal, ok := newString.Elem().(string)
	assert.True(t, ok)
	assert.Empty(t, stringVal)
}

func TestTypeOfStringPointer(t *testing.T) {
	ptrType := TypeOf[*string]()
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*string", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsString(typ))
	assert.Equal(t, "string", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	stringType, isString := ToString(typ)

	assert.NotNil(t, stringType)
	assert.True(t, isString)

	assert.False(t, stringType.CanSet())

	value, err := stringType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = stringType.SetValue("hello")
	assert.NotNil(t, err)

	newString := stringType.Instantiate()
	assert.NotNil(t, newString)

	stringPtrVal, ok := newString.Val().(*string)
	assert.True(t, ok)
	assert.Empty(t, *stringPtrVal)

	stringVal, ok := newString.Elem().(string)
	assert.True(t, ok)
	assert.Empty(t, stringVal)
}

func TestTypeOfStringObject(t *testing.T) {
	val := "hello"

	typ := TypeOfAny(val)
	assert.True(t, IsString(typ))
	assert.Equal(t, "string", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	stringType, isString := ToString(typ)

	assert.NotNil(t, stringType)
	assert.True(t, isString)

	assert.False(t, stringType.CanSet())

	value, err := stringType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)

	err = stringType.SetValue("world")
	assert.NotNil(t, err)
	assert.NotEqual(t, "world", val)

	newString := stringType.Instantiate()
	assert.NotNil(t, newString)

	stringPtrVal, ok := newString.Val().(*string)
	assert.True(t, ok)
	assert.Empty(t, *stringPtrVal)

	stringVal, ok := newString.Elem().(string)
	assert.True(t, ok)
	assert.Empty(t, stringVal)
}

func TestTypeOfStringObjectPointer(t *testing.T) {
	val := "hello"

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*string", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsString(typ))
	assert.Equal(t, "string", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	stringType, isString := ToString(typ)

	assert.NotNil(t, stringType)
	assert.True(t, isString)

	assert.True(t, stringType.CanSet())

	value, err := stringType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)

	err = stringType.SetValue("world")
	assert.Nil(t, err)
	assert.Equal(t, "world", val)

	newString := stringType.Instantiate()
	assert.NotNil(t, newString)

	stringPtrVal, ok := newString.Val().(*string)
	assert.True(t, ok)
	assert.Empty(t, *stringPtrVal)

	stringVal, ok := newString.Elem().(string)
	assert.True(t, ok)
	assert.Empty(t, stringVal)
}
