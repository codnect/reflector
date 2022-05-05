package reflector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeOfSlice(t *testing.T) {
	typ := TypeOf[[]int]()
	assert.True(t, IsSlice(typ))
	assert.Equal(t, "[]int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	sliceType, isSlice := ToSlice(typ)

	assert.NotNil(t, sliceType)
	assert.True(t, isSlice)

	assert.False(t, sliceType.CanSet())

	len, err := sliceType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	cap, err := sliceType.Cap()
	assert.Equal(t, -1, cap)
	assert.NotNil(t, err)

	elem, err := sliceType.Get(0)
	assert.Nil(t, elem)
	assert.NotNil(t, err)

	err = sliceType.Set(0, "anyValue")
	assert.NotNil(t, err)

	newSlice, err := sliceType.Slice(0, 0)
	assert.Empty(t, newSlice)
	assert.NotNil(t, err)

	count, err := sliceType.Copy(nil)
	assert.Equal(t, -1, count)
	assert.NotNil(t, err)

	count, err = sliceType.Copy(make([]int, 2))
	assert.Equal(t, -1, count)
	assert.NotNil(t, err)

	value, err := sliceType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = sliceType.SetValue([]int{2, 5})
	assert.NotNil(t, err)

	newSliceVal, _ := sliceType.Instantiate()
	assert.NotNil(t, newSliceVal)

	slicePtrVal, ok := newSliceVal.Val().(*[]int)
	assert.True(t, ok)
	assert.Empty(t, *slicePtrVal)
	assert.Equal(t, []int{}, *slicePtrVal)

	sliceVal, ok := newSliceVal.Elem().([]int)
	assert.True(t, ok)
	assert.Empty(t, sliceVal)
	assert.Equal(t, []int{}, sliceVal)
}

func TestTypeOfSlicePointer(t *testing.T) {
	ptrType := TypeOf[*[]int]()
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*[]int", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsSlice(typ))
	assert.Equal(t, "[]int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	sliceType, isSlice := ToSlice(typ)

	assert.NotNil(t, sliceType)
	assert.True(t, isSlice)

	assert.False(t, sliceType.CanSet())

	len, err := sliceType.Len()
	assert.Equal(t, -1, len)
	assert.NotNil(t, err)

	cap, err := sliceType.Cap()
	assert.Equal(t, -1, cap)
	assert.NotNil(t, err)

	elem, err := sliceType.Get(0)
	assert.Nil(t, elem)
	assert.NotNil(t, err)

	err = sliceType.Set(0, "anyValue")
	assert.NotNil(t, err)

	newSlice, err := sliceType.Slice(0, 0)
	assert.Empty(t, newSlice)
	assert.NotNil(t, err)

	count, err := sliceType.Copy(nil)
	assert.Equal(t, -1, count)
	assert.NotNil(t, err)

	count, err = sliceType.Copy(make([]int, 2))
	assert.Equal(t, -1, count)
	assert.NotNil(t, err)

	value, err := sliceType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	err = sliceType.SetValue([]int{2, 5})
	assert.NotNil(t, err)

	newSliceVal, _ := sliceType.Instantiate()
	assert.NotNil(t, newSliceVal)

	slicePtrVal, ok := newSliceVal.Val().(*[]int)
	assert.True(t, ok)
	assert.Empty(t, *slicePtrVal)
	assert.Equal(t, []int{}, *slicePtrVal)

	sliceVal, ok := newSliceVal.Elem().([]int)
	assert.True(t, ok)
	assert.Empty(t, sliceVal)
	assert.Equal(t, []int{}, sliceVal)
}

func TestTypeOfSliceObject(t *testing.T) {
	val := []int{5, 1, 8}

	typ := TypeOfAny(val)
	assert.True(t, IsSlice(typ))
	assert.Equal(t, "[]int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	sliceType, isSlice := ToSlice(typ)

	assert.NotNil(t, sliceType)
	assert.True(t, isSlice)

	assert.False(t, sliceType.CanSet())

	len, err := sliceType.Len()
	assert.Equal(t, 3, len)
	assert.Nil(t, err)

	cap, err := sliceType.Cap()
	assert.Equal(t, 3, cap)
	assert.Nil(t, err)

	elem, err := sliceType.Get(1)
	assert.Equal(t, 1, elem)
	assert.Nil(t, err)

	err = sliceType.Set(0, "anyValue")
	assert.NotNil(t, err)

	newSlice, err := sliceType.Slice(0, 0)
	assert.Empty(t, newSlice)
	assert.Nil(t, err)

	count, err := sliceType.Copy(nil)
	assert.Equal(t, -1, count)
	assert.NotNil(t, err)

	copySlice := make([]int, 2)
	count, err = sliceType.Copy(copySlice)
	assert.Equal(t, 2, count)
	assert.Nil(t, err)
	assert.Equal(t, []int{5, 1}, copySlice)

	value, err := sliceType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)
	assert.Equal(t, []int{5, 1, 8}, value)

	err = sliceType.SetValue([]int{2, 5})
	assert.NotNil(t, err)

	value, err = sliceType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)
	assert.Equal(t, []int{5, 1, 8}, value)

	newSliceVal, _ := sliceType.Instantiate()
	assert.NotNil(t, newSliceVal)

	slicePtrVal, ok := newSliceVal.Val().(*[]int)
	assert.True(t, ok)
	assert.Empty(t, *slicePtrVal)
	assert.Equal(t, []int{}, *slicePtrVal)

	sliceVal, ok := newSliceVal.Elem().([]int)
	assert.True(t, ok)
	assert.Empty(t, sliceVal)
	assert.Equal(t, []int{}, sliceVal)
}

func TestTypeOfSliceObjectPointer(t *testing.T) {
	val := []int{5, 1, 8}

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*[]int", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsSlice(typ))
	assert.Equal(t, "[]int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	sliceType, isSlice := ToSlice(typ)

	assert.NotNil(t, sliceType)
	assert.True(t, isSlice)

	assert.True(t, sliceType.CanSet())

	len, err := sliceType.Len()
	assert.Equal(t, 3, len)
	assert.Nil(t, err)

	cap, err := sliceType.Cap()
	assert.Equal(t, 3, cap)
	assert.Nil(t, err)

	elem, err := sliceType.Get(1)
	assert.Equal(t, 1, elem)
	assert.Nil(t, err)

	err = sliceType.Set(2, 3)
	assert.Nil(t, err)

	elem, err = sliceType.Get(2)
	assert.Equal(t, 3, elem)
	assert.Nil(t, err)

	newSlice, err := sliceType.Slice(0, 0)
	assert.Empty(t, newSlice)
	assert.Nil(t, err)

	count, err := sliceType.Copy(nil)
	assert.Equal(t, -1, count)
	assert.NotNil(t, err)

	copySlice := make([]int, 2)
	count, err = sliceType.Copy(copySlice)
	assert.Equal(t, 2, count)
	assert.Nil(t, err)
	assert.Equal(t, []int{5, 1}, copySlice)

	value, err := sliceType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)
	assert.Equal(t, []int{5, 1, 3}, value)

	err = sliceType.SetValue([]int{2, 5})
	assert.Nil(t, err)

	value, err = sliceType.Value()
	assert.NotEmpty(t, value)
	assert.Nil(t, err)
	assert.Equal(t, []int{2, 5}, value)

	len, err = sliceType.Len()
	assert.Equal(t, 2, len)
	assert.Nil(t, err)

	cap, err = sliceType.Cap()
	assert.Equal(t, 2, cap)
	assert.Nil(t, err)

	newSliceVal, _ := sliceType.Instantiate()
	assert.NotNil(t, newSliceVal)

	slicePtrVal, ok := newSliceVal.Val().(*[]int)
	assert.True(t, ok)
	assert.Empty(t, *slicePtrVal)
	assert.Equal(t, []int{}, *slicePtrVal)

	sliceVal, ok := newSliceVal.Elem().([]int)
	assert.True(t, ok)
	assert.Empty(t, sliceVal)
	assert.Equal(t, []int{}, sliceVal)
}
