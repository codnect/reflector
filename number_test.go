package reflector

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestTypeOfInt8(t *testing.T) {
	typ := TypeOf[int8]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int8", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize8, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int8]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int8]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	floatValue, err := signedInteger.Convert(TypeOf[float32]())
	assert.Nil(t, floatValue)
	assert.NotNil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(150))
	assert.False(t, signedInteger.Overflow(110))
	assert.False(t, signedInteger.Overflow(-127))
	assert.True(t, signedInteger.Overflow(-179))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int8)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int8)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt8Pointer(t *testing.T) {
	ptrType := TypeOf[*int8]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int8", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int8", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize8, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int8]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int8]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	floatValue, err := signedInteger.Convert(TypeOf[float32]())
	assert.Nil(t, floatValue)
	assert.NotNil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(150))
	assert.False(t, signedInteger.Overflow(110))
	assert.False(t, signedInteger.Overflow(-127))
	assert.True(t, signedInteger.Overflow(-179))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int8)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int8)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt8Object(t *testing.T) {
	val := int8(15)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int8", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize8, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int8]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int8]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	floatValue, err := signedInteger.Convert(TypeOf[float32]())
	assert.NotNil(t, floatValue)
	assert.Equal(t, float32(15.0), floatValue.Val())
	assert.Nil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(15), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(150))
	assert.False(t, signedInteger.Overflow(110))
	assert.False(t, signedInteger.Overflow(-127))
	assert.True(t, signedInteger.Overflow(-179))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int8)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int8)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt8ObjectPointer(t *testing.T) {
	val := int8(15)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int8", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int8", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize8, signedInteger.BitSize())

	assert.True(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int8]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int8]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	floatValue, err := signedInteger.Convert(TypeOf[float32]())
	assert.NotNil(t, floatValue)
	assert.Equal(t, float32(15.0), floatValue.Val())
	assert.Nil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(15), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(125)
	assert.Nil(t, err)
	assert.Equal(t, int8(125), val)

	err = signedInteger.SetValue(150)
	assert.Nil(t, err)
	assert.Equal(t, int8(-106), val)

	assert.True(t, signedInteger.Overflow(150))
	assert.False(t, signedInteger.Overflow(110))
	assert.False(t, signedInteger.Overflow(-127))
	assert.True(t, signedInteger.Overflow(-179))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int8)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int8)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt16(t *testing.T) {
	typ := TypeOf[int16]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int16", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize16, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int16]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int16]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	stringValue, err := signedInteger.Convert(TypeOf[string]())
	assert.Nil(t, stringValue)
	assert.NotNil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(33645))
	assert.False(t, signedInteger.Overflow(110))
	assert.False(t, signedInteger.Overflow(-127))
	assert.True(t, signedInteger.Overflow(-33632))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int16)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int16)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt16Pointer(t *testing.T) {
	ptrType := TypeOf[*int16]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int16", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int16", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize16, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int16]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int16]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	stringValue, err := signedInteger.Convert(TypeOf[string]())
	assert.Nil(t, stringValue)
	assert.NotNil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(33645))
	assert.False(t, signedInteger.Overflow(110))
	assert.False(t, signedInteger.Overflow(-127))
	assert.True(t, signedInteger.Overflow(-33632))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int16)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int16)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt16Object(t *testing.T) {
	val := int16(1024)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int16", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize16, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int16]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int16]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.NotNil(t, intValue)
	assert.Equal(t, 1024, intValue.Val())
	assert.Nil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(1024), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(33645))
	assert.False(t, signedInteger.Overflow(9733))
	assert.False(t, signedInteger.Overflow(-4546))
	assert.True(t, signedInteger.Overflow(-33632))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int16)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int16)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt16ObjectPointer(t *testing.T) {
	val := int16(1024)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int16", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int16", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize16, signedInteger.BitSize())

	assert.True(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int16]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int16]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.NotNil(t, intValue)
	assert.Equal(t, 1024, intValue.Val())
	assert.Nil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(1024), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(7032)
	assert.Nil(t, err)
	assert.Equal(t, int16(7032), val)

	err = signedInteger.SetValue(50235)
	assert.Nil(t, err)
	assert.Equal(t, int16(-15301), val)

	assert.True(t, signedInteger.Overflow(33645))
	assert.False(t, signedInteger.Overflow(9733))
	assert.False(t, signedInteger.Overflow(-4546))
	assert.True(t, signedInteger.Overflow(-33632))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int16)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int16)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt32(t *testing.T) {
	typ := TypeOf[int32]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize32, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int32]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int32]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.Nil(t, intValue)
	assert.NotNil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(2447483647))
	assert.False(t, signedInteger.Overflow(123643))
	assert.False(t, signedInteger.Overflow(-654756))
	assert.True(t, signedInteger.Overflow(-2447483647))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt32Pointer(t *testing.T) {
	ptrType := TypeOf[*int32]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int32", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize32, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int32]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int32]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.Nil(t, intValue)
	assert.NotNil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(2447483647))
	assert.False(t, signedInteger.Overflow(123643))
	assert.False(t, signedInteger.Overflow(-654756))
	assert.True(t, signedInteger.Overflow(-2447483647))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt32Object(t *testing.T) {
	val := int32(52123)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize32, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int32]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int32]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.NotNil(t, intValue)
	assert.Equal(t, 52123, intValue.Val())
	assert.Nil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(52123), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	assert.True(t, signedInteger.Overflow(2447483647))
	assert.False(t, signedInteger.Overflow(123643))
	assert.False(t, signedInteger.Overflow(-654756))
	assert.True(t, signedInteger.Overflow(-2447483647))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt32ObjectPointer(t *testing.T) {
	val := int32(1024)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int32", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize32, signedInteger.BitSize())

	assert.True(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int32]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int32]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.NotNil(t, intValue)
	assert.Equal(t, 1024, intValue.Val())
	assert.Nil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(1024), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(7032)
	assert.Nil(t, err)
	assert.Equal(t, int32(7032), val)

	err = signedInteger.SetValue(2447483647)
	assert.Nil(t, err)
	assert.Equal(t, int32(-1847483649), val)

	assert.True(t, signedInteger.Overflow(2447483647))
	assert.False(t, signedInteger.Overflow(123643))
	assert.False(t, signedInteger.Overflow(-654756))
	assert.True(t, signedInteger.Overflow(-2447483647))

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt64(t *testing.T) {
	typ := TypeOf[int64]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize64, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int64]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int64]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.Nil(t, intValue)
	assert.NotNil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt64Pointer(t *testing.T) {
	ptrType := TypeOf[*int64]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int64", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize64, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int64]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int64]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.Nil(t, intValue)
	assert.NotNil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt64Object(t *testing.T) {
	val := int64(52123)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize64, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int64]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int64]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.NotNil(t, intValue)
	assert.Equal(t, 52123, intValue.Val())
	assert.Nil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(52123), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt64ObjectPointer(t *testing.T) {
	val := int64(1024)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int64", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize64, signedInteger.BitSize())

	assert.True(t, signedInteger.CanSet())
	assert.True(t, signedInteger.IsInstantiable())
	assert.True(t, signedInteger.Compare(TypeOf[int64]()))
	assert.False(t, signedInteger.Compare(TypeOf[*int64]()))
	assert.False(t, signedInteger.Compare(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[string]()))
	assert.True(t, signedInteger.CanConvert(TypeOf[float32]()))

	intValue, err := signedInteger.Convert(TypeOf[int]())
	assert.NotNil(t, intValue)
	assert.Equal(t, 1024, intValue.Val())
	assert.Nil(t, err)

	value, err := signedInteger.Value()
	assert.Equal(t, int64(1024), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(7032)
	assert.Nil(t, err)
	assert.Equal(t, int64(7032), val)

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfInt(t *testing.T) {
	typ := TypeOf[int]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize64, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfIntPointer(t *testing.T) {
	ptrType := TypeOf[*int]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize64, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())

	value, err := signedInteger.Value()
	assert.Equal(t, int64(-1), value)
	assert.NotNil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfIntObject(t *testing.T) {
	val := 52123

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize64, signedInteger.BitSize())

	assert.False(t, signedInteger.CanSet())

	value, err := signedInteger.Value()
	assert.Equal(t, int64(52123), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(435)
	assert.NotNil(t, err)

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfIntObjectPointer(t *testing.T) {
	val := 1024

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*int", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsSignedInteger(typ))
	assert.Equal(t, "int", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	signedInteger := ToSignedInteger(typ)

	assert.NotNil(t, signedInteger)
	assert.Equal(t, BitSize64, signedInteger.BitSize())

	assert.True(t, signedInteger.CanSet())

	value, err := signedInteger.Value()
	assert.Equal(t, int64(1024), value)
	assert.Nil(t, err)

	err = signedInteger.SetValue(7032)
	assert.Nil(t, err)
	assert.Equal(t, int(7032), val)

	newInteger, _ := signedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*int)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(int)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint8(t *testing.T) {
	typ := TypeOf[uint8]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint8", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize8, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(532))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))
	assert.True(t, unsignedInteger.Overflow(7653))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint8)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint8)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint8Pointer(t *testing.T) {
	ptrType := TypeOf[*uint8]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint8", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint8", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize8, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(532))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))
	assert.True(t, unsignedInteger.Overflow(7653))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint8)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint8)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint8Object(t *testing.T) {
	val := uint8(15)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint8", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize8, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(532))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))
	assert.True(t, unsignedInteger.Overflow(7653))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint8)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint8)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint8ObjectPointer(t *testing.T) {
	val := uint8(15)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint8", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint8", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize8, unsignedInteger.BitSize())

	assert.True(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.Nil(t, err)
	assert.Equal(t, uint8(125), val)

	err = unsignedInteger.SetValue(363)
	assert.Nil(t, err)
	assert.Equal(t, uint8(0x6b), val)

	assert.True(t, unsignedInteger.Overflow(532))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))
	assert.True(t, unsignedInteger.Overflow(7653))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint8)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint8)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint16(t *testing.T) {
	typ := TypeOf[uint16]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint16", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize16, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(73412))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))
	assert.True(t, unsignedInteger.Overflow(123535))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint16)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint16)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint16Pointer(t *testing.T) {
	ptrType := TypeOf[*uint16]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint16", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint16", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize16, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(73412))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))
	assert.True(t, unsignedInteger.Overflow(123535))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint16)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint16)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint16Object(t *testing.T) {
	val := uint16(15)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint16", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize16, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(73412))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))
	assert.True(t, unsignedInteger.Overflow(123535))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint16)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint16)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint16ObjectPointer(t *testing.T) {
	val := uint16(15)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint16", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint16", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize16, unsignedInteger.BitSize())

	assert.True(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.Nil(t, err)
	assert.Equal(t, uint16(125), val)

	err = unsignedInteger.SetValue(70321)
	assert.Nil(t, err)
	assert.Equal(t, uint16(0x12b1), val)

	assert.True(t, unsignedInteger.Overflow(73412))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))
	assert.True(t, unsignedInteger.Overflow(123535))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint16)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint16)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint32(t *testing.T) {
	typ := TypeOf[uint32]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize32, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(43213253123))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint32Pointer(t *testing.T) {
	ptrType := TypeOf[*uint32]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint32", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize32, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(43213253123))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint32Object(t *testing.T) {
	val := uint32(15)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize32, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, unsignedInteger.Overflow(43213253123))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint32ObjectPointer(t *testing.T) {
	val := uint32(15)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint32", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize32, unsignedInteger.BitSize())

	assert.True(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.Nil(t, err)
	assert.Equal(t, uint32(125), val)

	err = unsignedInteger.SetValue(43213253123)
	assert.Nil(t, err)
	assert.Equal(t, uint32(0xfb5ea03), val)

	assert.True(t, unsignedInteger.Overflow(43213253123))
	assert.False(t, unsignedInteger.Overflow(110))
	assert.False(t, unsignedInteger.Overflow(212))

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint64(t *testing.T) {
	typ := TypeOf[uint64]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize64, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint64Pointer(t *testing.T) {
	ptrType := TypeOf[*uint64]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint64", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize64, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint64Object(t *testing.T) {
	val := uint64(15)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize64, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint64ObjectPointer(t *testing.T) {
	val := uint64(15)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint64", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize64, unsignedInteger.BitSize())

	assert.True(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.Nil(t, err)
	assert.Equal(t, uint64(125), val)

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUint(t *testing.T) {
	typ := TypeOf[uint]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize64, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUintPointer(t *testing.T) {
	ptrType := TypeOf[*uint]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize64, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(0), value)
	assert.NotNil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUintObject(t *testing.T) {
	val := uint(15)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize64, unsignedInteger.BitSize())

	assert.False(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfUintObjectPointer(t *testing.T) {
	val := uint(15)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*uint", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsInteger(typ))
	assert.True(t, IsUnsignedInteger(typ))
	assert.Equal(t, "uint", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	unsignedInteger := ToUnsignedInteger(typ)

	assert.NotNil(t, unsignedInteger)
	assert.Equal(t, BitSize64, unsignedInteger.BitSize())

	assert.True(t, unsignedInteger.CanSet())

	value, err := unsignedInteger.Value()
	assert.Equal(t, uint64(15), value)
	assert.Nil(t, err)

	err = unsignedInteger.SetValue(125)
	assert.Nil(t, err)
	assert.Equal(t, uint(125), val)

	newInteger, _ := unsignedInteger.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*uint)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(uint)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfFloat32(t *testing.T) {
	typ := TypeOf[float32]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsFloat(typ))
	assert.Equal(t, "float32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	float := ToFloat(typ)

	assert.NotNil(t, float)
	assert.Equal(t, BitSize32, float.BitSize())

	assert.False(t, float.CanSet())

	value, err := float.Value()
	assert.Equal(t, float64(0), value)
	assert.NotNil(t, err)

	err = float.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, float.Overflow(math.MaxFloat64))
	assert.False(t, float.Overflow(110))
	assert.False(t, float.Overflow(-127))

	newInteger, _ := float.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*float32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(float32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfFloat32Pointer(t *testing.T) {
	ptrType := TypeOf[*float32]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*float32", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsFloat(typ))
	assert.Equal(t, "float32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	float := ToFloat(typ)

	assert.NotNil(t, float)
	assert.Equal(t, BitSize32, float.BitSize())

	assert.False(t, float.CanSet())

	value, err := float.Value()
	assert.Equal(t, float64(0), value)
	assert.NotNil(t, err)

	err = float.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, float.Overflow(math.MaxFloat64))
	assert.False(t, float.Overflow(110))
	assert.False(t, float.Overflow(-127))

	newInteger, _ := float.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*float32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(float32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfFloat32Object(t *testing.T) {
	val := float32(15)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsFloat(typ))
	assert.Equal(t, "float32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	float := ToFloat(typ)

	assert.NotNil(t, float)
	assert.Equal(t, BitSize32, float.BitSize())

	assert.False(t, float.CanSet())

	value, err := float.Value()
	assert.Equal(t, float64(15), value)
	assert.Nil(t, err)

	err = float.SetValue(125)
	assert.NotNil(t, err)

	assert.True(t, float.Overflow(math.MaxFloat64))
	assert.False(t, float.Overflow(110))
	assert.False(t, float.Overflow(-127))

	newInteger, _ := float.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*float32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(float32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfFloat32ObjectPointer(t *testing.T) {
	val := float32(15)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*float32", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsFloat(typ))
	assert.Equal(t, "float32", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	float := ToFloat(typ)

	assert.NotNil(t, float)
	assert.Equal(t, BitSize32, float.BitSize())

	assert.True(t, float.CanSet())

	value, err := float.Value()
	assert.Equal(t, float64(15), value)
	assert.Nil(t, err)

	err = float.SetValue(125)
	assert.Nil(t, err)

	assert.True(t, float.Overflow(math.MaxFloat64))
	assert.False(t, float.Overflow(110))
	assert.False(t, float.Overflow(-127))

	newInteger, _ := float.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*float32)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(float32)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfFloat64(t *testing.T) {
	typ := TypeOf[float64]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsFloat(typ))
	assert.Equal(t, "float64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	float := ToFloat(typ)

	assert.NotNil(t, float)
	assert.Equal(t, BitSize64, float.BitSize())

	assert.False(t, float.CanSet())

	value, err := float.Value()
	assert.Equal(t, float64(0), value)
	assert.NotNil(t, err)

	err = float.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := float.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*float64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(float64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfFloat64Pointer(t *testing.T) {
	ptrType := TypeOf[*float64]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*float64", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsFloat(typ))
	assert.Equal(t, "float64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	float := ToFloat(typ)

	assert.NotNil(t, float)
	assert.Equal(t, BitSize64, float.BitSize())

	assert.False(t, float.CanSet())

	value, err := float.Value()
	assert.Equal(t, float64(0), value)
	assert.NotNil(t, err)

	err = float.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := float.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*float64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(float64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfFloat64Object(t *testing.T) {
	val := float64(15)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsFloat(typ))
	assert.Equal(t, "float64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	float := ToFloat(typ)

	assert.NotNil(t, float)
	assert.Equal(t, BitSize64, float.BitSize())

	assert.False(t, float.CanSet())

	value, err := float.Value()
	assert.Equal(t, float64(15), value)
	assert.Nil(t, err)

	err = float.SetValue(125)
	assert.NotNil(t, err)

	newInteger, _ := float.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*float64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(float64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfFloat64ObjectPointer(t *testing.T) {
	val := float64(15)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*float64", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsFloat(typ))
	assert.Equal(t, "float64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	float := ToFloat(typ)

	assert.NotNil(t, float)
	assert.Equal(t, BitSize64, float.BitSize())

	assert.True(t, float.CanSet())

	value, err := float.Value()
	assert.Equal(t, float64(15), value)
	assert.Nil(t, err)

	err = float.SetValue(float64(125))
	assert.Nil(t, err)

	newInteger, _ := float.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*float64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(float64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfComplex64(t *testing.T) {
	typ := TypeOf[complex64]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsComplex(typ))
	assert.Equal(t, "complex64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	complexType := ToComplex(typ)

	assert.NotNil(t, complexType)
	assert.Equal(t, BitSize64, complexType.BitSize())

	assert.False(t, complexType.CanSet())

	imaginaryData, err := complexType.ImaginaryData()
	assert.Equal(t, float64(0), imaginaryData)
	assert.NotNil(t, err)

	err = complexType.SetImaginaryData(125)
	assert.NotNil(t, err)

	value, err := complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	realData, err := complexType.RealData()
	assert.Equal(t, float64(0), realData)
	assert.NotNil(t, err)

	err = complexType.SetRealData(23)
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	err = complexType.SetValue(complex(1, 2))
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	newInteger, _ := complexType.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*complex64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(complex64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfComplex64Pointer(t *testing.T) {
	ptrType := TypeOf[*complex64]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*complex64", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsComplex(typ))
	assert.Equal(t, "complex64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	complexType := ToComplex(typ)

	assert.NotNil(t, complexType)
	assert.Equal(t, BitSize64, complexType.BitSize())

	assert.False(t, complexType.CanSet())

	imaginaryData, err := complexType.ImaginaryData()
	assert.Equal(t, float64(0), imaginaryData)
	assert.NotNil(t, err)

	err = complexType.SetImaginaryData(125)
	assert.NotNil(t, err)

	value, err := complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	realData, err := complexType.RealData()
	assert.Equal(t, float64(0), realData)
	assert.NotNil(t, err)

	err = complexType.SetRealData(23)
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	err = complexType.SetValue(complex(1, 2))
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	newInteger, _ := complexType.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*complex64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(complex64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfComplex64Object(t *testing.T) {
	val := complex64(complex(15, 2))

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsComplex(typ))
	assert.Equal(t, "complex64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	complexType := ToComplex(typ)

	assert.NotNil(t, complexType)
	assert.Equal(t, BitSize64, complexType.BitSize())

	assert.False(t, complexType.CanSet())

	imaginaryData, err := complexType.ImaginaryData()
	assert.Equal(t, float64(2), imaginaryData)
	assert.Nil(t, err)

	err = complexType.SetImaginaryData(125)
	assert.NotNil(t, err)

	value, err := complexType.Value()
	assert.Equal(t, complex(15, 2), value)
	assert.Nil(t, err)

	realData, err := complexType.RealData()
	assert.Equal(t, float64(15), realData)
	assert.Nil(t, err)

	err = complexType.SetRealData(23)
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(15, 2), value)
	assert.Nil(t, err)

	err = complexType.SetValue(complex(1, 2))
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(15, 2), value)
	assert.Nil(t, err)

	newInteger, _ := complexType.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*complex64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(complex64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfComplex64ObjectPointer(t *testing.T) {
	val := complex64(complex(15, 2))

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*complex64", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsComplex(typ))
	assert.Equal(t, "complex64", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	complexType := ToComplex(typ)

	assert.NotNil(t, complexType)
	assert.Equal(t, BitSize64, complexType.BitSize())

	assert.True(t, complexType.CanSet())

	imaginaryData, err := complexType.ImaginaryData()
	assert.Equal(t, float64(2), imaginaryData)
	assert.Nil(t, err)

	err = complexType.SetImaginaryData(125)
	assert.Nil(t, err)

	value, err := complexType.Value()
	assert.Equal(t, complex(15, 125), value)
	assert.Nil(t, err)

	realData, err := complexType.RealData()
	assert.Equal(t, float64(15), realData)
	assert.Nil(t, err)

	err = complexType.SetRealData(23)
	assert.Nil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(23, 125), value)
	assert.Nil(t, err)

	err = complexType.SetValue(complex(1, 2))
	assert.Nil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(1, 2), value)
	assert.Nil(t, err)

	newInteger, _ := complexType.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*complex64)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(complex64)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfComplex128(t *testing.T) {
	typ := TypeOf[complex128]()
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsComplex(typ))
	assert.Equal(t, "complex128", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	complexType := ToComplex(typ)

	assert.NotNil(t, complexType)
	assert.Equal(t, BitSize128, complexType.BitSize())

	assert.False(t, complexType.CanSet())

	imaginaryData, err := complexType.ImaginaryData()
	assert.Equal(t, float64(0), imaginaryData)
	assert.NotNil(t, err)

	err = complexType.SetImaginaryData(125)
	assert.NotNil(t, err)

	value, err := complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	realData, err := complexType.RealData()
	assert.Equal(t, float64(0), realData)
	assert.NotNil(t, err)

	err = complexType.SetRealData(23)
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	err = complexType.SetValue(complex(1, 2))
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	newInteger, _ := complexType.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*complex128)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(complex128)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfComplex128Pointer(t *testing.T) {
	ptrType := TypeOf[*complex128]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*complex128", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsComplex(typ))
	assert.Equal(t, "complex128", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	complexType := ToComplex(typ)

	assert.NotNil(t, complexType)
	assert.Equal(t, BitSize128, complexType.BitSize())

	assert.False(t, complexType.CanSet())

	imaginaryData, err := complexType.ImaginaryData()
	assert.Equal(t, float64(0), imaginaryData)
	assert.NotNil(t, err)

	err = complexType.SetImaginaryData(125)
	assert.NotNil(t, err)

	value, err := complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	realData, err := complexType.RealData()
	assert.Equal(t, float64(0), realData)
	assert.NotNil(t, err)

	err = complexType.SetRealData(23)
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	err = complexType.SetValue(complex(1, 2))
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(0, 0), value)
	assert.NotNil(t, err)

	newInteger, _ := complexType.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*complex128)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(complex128)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfComplex128Object(t *testing.T) {
	val := complex(15, 2)

	typ := TypeOfAny(val)
	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsComplex(typ))
	assert.Equal(t, "complex128", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	complexType := ToComplex(typ)

	assert.NotNil(t, complexType)
	assert.Equal(t, BitSize128, complexType.BitSize())

	assert.False(t, complexType.CanSet())

	imaginaryData, err := complexType.ImaginaryData()
	assert.Equal(t, float64(2), imaginaryData)
	assert.Nil(t, err)

	err = complexType.SetImaginaryData(125)
	assert.NotNil(t, err)

	value, err := complexType.Value()
	assert.Equal(t, complex(15, 2), value)
	assert.Nil(t, err)

	realData, err := complexType.RealData()
	assert.Equal(t, float64(15), realData)
	assert.Nil(t, err)

	err = complexType.SetRealData(23)
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(15, 2), value)
	assert.Nil(t, err)

	err = complexType.SetValue(complex(1, 2))
	assert.NotNil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(15, 2), value)
	assert.Nil(t, err)

	newInteger, _ := complexType.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*complex128)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(complex128)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}

func TestTypeOfComplex128ObjectPointer(t *testing.T) {
	val := complex(15, 2)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*complex128", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsBasic(typ))
	assert.True(t, IsNumber(typ))
	assert.True(t, IsComplex(typ))
	assert.Equal(t, "complex128", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	complexType := ToComplex(typ)

	assert.NotNil(t, complexType)
	assert.Equal(t, BitSize128, complexType.BitSize())

	assert.True(t, complexType.CanSet())

	imaginaryData, err := complexType.ImaginaryData()
	assert.Equal(t, float64(2), imaginaryData)
	assert.Nil(t, err)

	err = complexType.SetImaginaryData(125)
	assert.Nil(t, err)

	value, err := complexType.Value()
	assert.Equal(t, complex(15, 125), value)
	assert.Nil(t, err)

	realData, err := complexType.RealData()
	assert.Equal(t, float64(15), realData)
	assert.Nil(t, err)

	err = complexType.SetRealData(23)
	assert.Nil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(23, 125), value)
	assert.Nil(t, err)

	err = complexType.SetValue(complex(1, 2))
	assert.Nil(t, err)

	value, err = complexType.Value()
	assert.Equal(t, complex(1, 2), value)
	assert.Nil(t, err)

	newInteger, _ := complexType.Instantiate()
	assert.NotNil(t, newInteger)

	integerPtrVal, ok := newInteger.Val().(*complex128)
	assert.True(t, ok)
	assert.Empty(t, *integerPtrVal)

	integerVal, ok := newInteger.Elem().(complex128)
	assert.True(t, ok)
	assert.Empty(t, integerVal)
}
