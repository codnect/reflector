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
	assert.True(t, stringType.IsInstantiable())
	assert.False(t, stringType.Compare(TypeOf[bool]()))
	assert.False(t, stringType.Compare(TypeOf[*bool]()))
	assert.True(t, stringType.Compare(TypeOf[string]()))

	value, err := stringType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = stringType.SetValue("anyTestValue")
	assert.NotNil(t, err)

	stringValue, err := stringType.StringValue()
	assert.Empty(t, stringValue)
	assert.NotNil(t, err)

	err = stringType.SetStringValue("anyTestValue")
	assert.NotNil(t, err)

	newString, _ := stringType.Instantiate()
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

	assert.False(t, ptrType.CanSet())
	assert.True(t, ptrType.IsInstantiable())
	assert.False(t, ptrType.Compare(TypeOf[string]()))
	assert.True(t, ptrType.Compare(TypeOf[*string]()))
	assert.False(t, ptrType.Compare(TypeOf[bool]()))

	value, err := ptrType.Value()
	assert.Nil(t, value)
	assert.NotNil(t, err)

	err = ptrType.SetValue(nil)
	assert.NotNil(t, err)

	value, err = ptrType.Value()
	assert.Nil(t, value)
	assert.NotNil(t, err)

	newString, _ := ptrType.Instantiate()
	assert.NotNil(t, newString)

	stringPtrVal, ok := newString.Val().(*string)
	assert.True(t, ok)
	assert.Empty(t, *stringPtrVal)

	stringVal, ok := newString.Elem().(string)
	assert.True(t, ok)
	assert.Empty(t, stringVal)

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
	assert.True(t, stringType.IsInstantiable())
	assert.False(t, stringType.Compare(TypeOf[bool]()))
	assert.False(t, stringType.Compare(TypeOf[*bool]()))
	assert.True(t, stringType.Compare(TypeOf[string]()))

	value, err = stringType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = stringType.SetValue("anyTestValue")
	assert.NotNil(t, err)

	stringValue, err := stringType.StringValue()
	assert.Empty(t, stringValue)
	assert.NotNil(t, err)

	err = stringType.SetStringValue("anyTestValue")
	assert.NotNil(t, err)

	newString, _ = stringType.Instantiate()
	assert.NotNil(t, newString)

	stringPtrVal, ok = newString.Val().(*string)
	assert.True(t, ok)
	assert.Empty(t, *stringPtrVal)

	stringVal, ok = newString.Elem().(string)
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
	assert.True(t, stringType.IsInstantiable())
	assert.False(t, stringType.Compare(TypeOf[bool]()))
	assert.False(t, stringType.Compare(TypeOf[*bool]()))
	assert.True(t, stringType.Compare(TypeOf[string]()))

	value, err := stringType.Value()
	assert.NotEmpty(t, value)
	assert.Equal(t, val, value)
	assert.Nil(t, err)

	err = stringType.SetValue("anyTestValue")
	assert.NotNil(t, err)

	stringValue, err := stringType.StringValue()
	assert.NotNil(t, stringValue)
	assert.Equal(t, val, value)
	assert.Nil(t, err)

	err = stringType.SetStringValue("anyTestValue")
	assert.NotNil(t, err)

	newString, _ := stringType.Instantiate()
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

	assert.False(t, ptrType.CanSet())
	assert.True(t, ptrType.IsInstantiable())
	assert.False(t, ptrType.Compare(TypeOf[string]()))
	assert.True(t, ptrType.Compare(TypeOf[*string]()))
	assert.False(t, ptrType.Compare(TypeOf[bool]()))

	value, err := ptrType.Value()
	assert.NotEmpty(t, value)
	assert.Equal(t, &val, value)
	assert.Nil(t, err)

	err = ptrType.SetValue("anyTestValue")
	assert.NotNil(t, err)

	value, err = ptrType.Value()
	assert.NotEmpty(t, value)
	assert.Equal(t, &val, value)
	assert.Nil(t, err)

	newString, _ := ptrType.Instantiate()
	assert.NotNil(t, newString)

	stringPtrVal, ok := newString.Val().(*string)
	assert.True(t, ok)
	assert.Empty(t, *stringPtrVal)

	stringVal, ok := newString.Elem().(string)
	assert.True(t, ok)
	assert.Empty(t, stringVal)

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
	assert.True(t, stringType.IsInstantiable())
	assert.False(t, stringType.Compare(TypeOf[bool]()))
	assert.False(t, stringType.Compare(TypeOf[*bool]()))
	assert.True(t, stringType.Compare(TypeOf[string]()))

	value, err = stringType.Value()
	assert.NotEmpty(t, value)
	assert.Equal(t, val, value)
	assert.Nil(t, err)

	err = stringType.SetValue("anyTestValue")
	assert.Nil(t, err)

	stringValue, err := stringType.StringValue()
	assert.NotEmpty(t, stringValue)
	assert.Equal(t, "anyTestValue", stringValue)
	assert.Nil(t, err)

	err = stringType.SetStringValue("anyTestValue2")
	assert.Nil(t, err)

	newString, _ = stringType.Instantiate()
	assert.NotNil(t, newString)

	stringPtrVal, ok = newString.Val().(*string)
	assert.True(t, ok)
	assert.Empty(t, *stringPtrVal)

	stringVal, ok = newString.Elem().(string)
	assert.True(t, ok)
	assert.Empty(t, stringVal)
}
