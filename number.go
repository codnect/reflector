package reflector

import (
	"math/bits"
	"reflect"
)

type BitSize int

const (
	BitSize8   BitSize = 8
	BitSize16  BitSize = 16
	BitSize32  BitSize = 32
	BitSize64  BitSize = 64
	BitSize128 BitSize = 128
)

func bitSize(kind reflect.Kind) BitSize {
	switch kind {
	case reflect.Int:
		if bits.UintSize == 32 {
			return BitSize32
		}
		return BitSize64
	case reflect.Int8:
		return BitSize8
	case reflect.Int16:
		return BitSize16
	case reflect.Int32:
		return BitSize32
	case reflect.Int64:
		return BitSize64
	case reflect.Uint:
		if bits.UintSize == 32 {
			return BitSize32
		}
		return BitSize64
	case reflect.Uint8:
		return BitSize8
	case reflect.Uint16:
		return BitSize16
	case reflect.Uint32:
		return BitSize32
	case reflect.Uint64:
		return BitSize64
	case reflect.Float32:
		return BitSize32
	case reflect.Float64:
		return BitSize64
	case reflect.Complex64:
		return BitSize64
	case reflect.Complex128:
		return BitSize128
	}

	panic("Invalid kind")
}

type SignedInteger interface {
	Type
	Instantiable
	BitSize() BitSize
	Value() int64
	SetValue(v int64)
	Overflow(v int64) bool
}

type signedInteger struct {
	bitSize BitSize

	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *signedInteger) Name() string {
	return s.reflectType.Name()
}

func (s *signedInteger) PackageName() string {
	return s.reflectType.Name()
}

func (s *signedInteger) HasReference() bool {
	return s.reflectValue != nil
}

func (s *signedInteger) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *signedInteger) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *signedInteger) BitSize() BitSize {
	return s.bitSize
}

func (s *signedInteger) Value() int64 {
	return 0
}

func (s *signedInteger) SetValue(v int64) {

}

func (s *signedInteger) Overflow(v int64) bool {
	return false
}

func (s *signedInteger) Instantiate() any {
	return reflect.New(s.reflectType).Interface()
}

type UnsignedInteger interface {
	Type
	Instantiable
	BitSize() BitSize
	Value() uint64
	SetValue(v uint64)
	Overflow(v uint64) bool
}

type unsignedInteger struct {
	bitSize BitSize

	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *unsignedInteger) Name() string {
	return s.reflectType.Name()
}

func (s *unsignedInteger) PackageName() string {
	return s.reflectType.Name()
}

func (s *unsignedInteger) HasReference() bool {
	return s.reflectValue != nil
}

func (s *unsignedInteger) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *unsignedInteger) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *unsignedInteger) BitSize() BitSize {
	return s.bitSize
}

func (s *unsignedInteger) Value() uint64 {
	return 0
}

func (s *unsignedInteger) SetValue(v uint64) {

}

func (s *unsignedInteger) Overflow(v uint64) bool {
	return false
}

func (s *unsignedInteger) Instantiate() any {
	return reflect.New(s.reflectType).Interface()
}

type Float interface {
	Type
	Instantiable
	BitSize() BitSize
	Value() float64
	SetValue(v float64)
	Overflow(v float64) bool
}

type float struct {
	bitSize BitSize

	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (f *float) Name() string {
	return f.reflectType.Name()
}

func (f *float) PackageName() string {
	return f.reflectType.Name()
}

func (f *float) HasReference() bool {
	return f.reflectValue != nil
}

func (f *float) ReflectType() reflect.Type {
	return f.reflectType
}

func (f *float) ReflectValue() *reflect.Value {
	return f.reflectValue
}

func (f *float) BitSize() BitSize {
	return f.bitSize
}

func (f *float) Value() float64 {
	return 0
}

func (f *float) SetValue(v float64) {

}

func (f *float) Overflow(v float64) bool {
	return false
}

func (f *float) Instantiate() any {
	return reflect.New(f.reflectType).Interface()
}

type Complex interface {
	Type
	Instantiable
	BitSize() BitSize
	ImaginaryData() complex128
	RealData() complex128
	SetImaginaryData(val complex128)
	SetRealData(val complex128)
}

type complex struct {
	bitSize BitSize

	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *complex) Name() string {
	return s.reflectType.Name()
}

func (s *complex) PackageName() string {
	return s.reflectType.Name()
}

func (s *complex) HasReference() bool {
	return s.reflectValue != nil
}

func (s *complex) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *complex) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *complex) BitSize() BitSize {
	return s.bitSize
}

func (s *complex) ImaginaryData() complex128 {
	return 0
}

func (s *complex) RealData() complex128 {
	return 0
}

func (s *complex) SetImaginaryData(val complex128) {

}

func (s *complex) SetRealData(val complex128) {

}

func (s *complex) Instantiate() any {
	return reflect.New(s.reflectType).Interface()
}
