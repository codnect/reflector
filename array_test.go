package reflector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeOfArray(t *testing.T) {
	typ := TypeOf[[2]int]()
	assert.True(t, IsArray(typ))
	assert.Equal(t, "[2]int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	arrayType, isArray := ToArray(typ)

	assert.NotNil(t, arrayType)
	assert.True(t, isArray)

	assert.False(t, arrayType.CanSet())

	value, err := arrayType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = arrayType.SetValue([2]int{2, 5})
	assert.NotNil(t, err)

	newArray := arrayType.Instantiate()
	assert.NotNil(t, newArray)

	arrayPtrVal, ok := newArray.Val().(*[2]int)
	assert.True(t, ok)
	assert.NotEmpty(t, *arrayPtrVal)
	assert.Equal(t, [2]int{0, 0}, *arrayPtrVal)

	arrayVal, ok := newArray.Elem().([2]int)
	assert.True(t, ok)
	assert.NotEmpty(t, arrayVal)
	assert.Equal(t, [2]int{0, 0}, arrayVal)
}

func TestTypeOfArrayPointer(t *testing.T) {
	ptrType := TypeOf[*[2]int]()
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*[2]int", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsArray(typ))
	assert.Equal(t, "[2]int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	arrayType, isArray := ToArray(typ)

	assert.NotNil(t, arrayType)
	assert.True(t, isArray)

	assert.False(t, arrayType.CanSet())

	value, err := arrayType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = arrayType.SetValue([2]int{2, 5})
	assert.NotNil(t, err)

	newArray := arrayType.Instantiate()
	assert.NotNil(t, newArray)

	arrayPtrVal, ok := newArray.Val().(*[2]int)
	assert.True(t, ok)
	assert.NotEmpty(t, *arrayPtrVal)
	assert.Equal(t, [2]int{0, 0}, *arrayPtrVal)

	arrayVal, ok := newArray.Elem().([2]int)
	assert.True(t, ok)
	assert.NotEmpty(t, arrayVal)
	assert.Equal(t, [2]int{0, 0}, arrayVal)
}

func TestTypeOfArrayObject(t *testing.T) {
	val := [3]int{5, 1, 8}

	typ := TypeOfAny(val)
	assert.True(t, IsArray(typ))
	assert.Equal(t, "[3]int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	arrayType, isArray := ToArray(typ)

	assert.NotNil(t, arrayType)
	assert.True(t, isArray)

	assert.False(t, arrayType.CanSet())

	value, err := arrayType.Value()
	assert.NotEmpty(t, value)
	assert.Equal(t, [3]int{5, 1, 8}, value)
	assert.Nil(t, err)

	err = arrayType.SetValue([2]int{2, 5})
	assert.NotNil(t, err)
	assert.NotEqual(t, [2]int{2, 5}, val)

	len := arrayType.Len()
	assert.Equal(t, 3, len)

	item, err := arrayType.Get(-1)
	assert.Nil(t, item)
	assert.NotNil(t, err)

	item, err = arrayType.Get(0)
	assert.Equal(t, 5, item)
	assert.Nil(t, err)

	item, err = arrayType.Get(1)
	assert.Equal(t, 1, item)
	assert.Nil(t, err)

	item, err = arrayType.Get(2)
	assert.Equal(t, 8, item)
	assert.Nil(t, err)

	item, err = arrayType.Get(3)
	assert.Nil(t, item)
	assert.NotNil(t, err)

	err = arrayType.Set(-1, 2)
	assert.Nil(t, item)
	assert.NotNil(t, err)

	err = arrayType.Set(2, 9)
	assert.Nil(t, item)
	assert.NotNil(t, err)

	err = arrayType.Set(3, 2)
	assert.Nil(t, item)
	assert.NotNil(t, err)

	newArray := arrayType.Instantiate()
	assert.NotNil(t, newArray)

	arrayPtrVal, ok := newArray.Val().(*[3]int)
	assert.True(t, ok)
	assert.NotEmpty(t, *arrayPtrVal)
	assert.Equal(t, [3]int{0, 0, 0}, *arrayPtrVal)

	arrayVal, ok := newArray.Elem().([3]int)
	assert.True(t, ok)
	assert.NotEmpty(t, arrayVal)
	assert.Equal(t, [3]int{0, 0, 0}, arrayVal)
}

func TestTypeOfArrayObjectPointer(t *testing.T) {
	val := [3]int{5, 1, 8}

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*[3]int", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsArray(typ))
	assert.Equal(t, "[3]int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	arrayType, isArray := ToArray(typ)

	assert.NotNil(t, arrayType)
	assert.True(t, isArray)

	assert.True(t, arrayType.CanSet())

	value, err := arrayType.Value()
	assert.NotEmpty(t, value)
	assert.Equal(t, [3]int{5, 1, 8}, value)
	assert.Nil(t, err)

	err = arrayType.SetValue([3]int{2, 0, 6})
	assert.Nil(t, err)
	assert.Equal(t, [3]int{2, 0, 6}, val)

	len := arrayType.Len()
	assert.Equal(t, 3, len)

	item, err := arrayType.Get(-1)
	assert.Nil(t, item)
	assert.NotNil(t, err)

	item, err = arrayType.Get(0)
	assert.Equal(t, 2, item)
	assert.Nil(t, err)

	item, err = arrayType.Get(1)
	assert.Equal(t, 0, item)
	assert.Nil(t, err)

	item, err = arrayType.Get(2)
	assert.Equal(t, 6, item)
	assert.Nil(t, err)

	item, err = arrayType.Get(3)
	assert.Nil(t, item)
	assert.NotNil(t, err)

	err = arrayType.Set(-1, 2)
	assert.NotNil(t, err)

	err = arrayType.Set(2, 9)
	assert.Nil(t, err)

	item, err = arrayType.Get(2)
	assert.Equal(t, 9, item)
	assert.Nil(t, err)

	err = arrayType.Set(3, 2)
	assert.NotNil(t, err)

	newArray := arrayType.Instantiate()
	assert.NotNil(t, newArray)

	arrayPtrVal, ok := newArray.Val().(*[3]int)
	assert.True(t, ok)
	assert.NotEmpty(t, *arrayPtrVal)
	assert.Equal(t, [3]int{0, 0, 0}, *arrayPtrVal)

	arrayVal, ok := newArray.Elem().([3]int)

	assert.True(t, ok)
	assert.NotEmpty(t, arrayVal)
	assert.Equal(t, [3]int{0, 0, 0}, arrayVal)

	//	arrayType.Copy(arrayVal)
}
