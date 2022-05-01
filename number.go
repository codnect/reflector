package reflector

import (
	"errors"
	"math"
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
	CanSet() bool
	Value() (int64, error)
	SetValue(v int64) error
	Overflow(v int64) bool
}

type signedIntegerType struct {
	bitSize BitSize

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *signedIntegerType) Name() string {
	return s.reflectType.Name()
}

func (s *signedIntegerType) PackageName() string {
	return ""
}

func (s *signedIntegerType) PackagePath() string {
	return ""
}

func (s *signedIntegerType) HasValue() bool {
	return s.reflectValue != nil
}

func (s *signedIntegerType) Parent() Type {
	return s.parent
}

func (s *signedIntegerType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *signedIntegerType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *signedIntegerType) BitSize() BitSize {
	return s.bitSize
}

func (s *signedIntegerType) CanSet() bool {
	if s.reflectValue == nil {
		return false
	}

	return s.reflectValue.CanSet()
}

func (s *signedIntegerType) Value() (int64, error) {
	if s.reflectValue == nil {
		return -1, errors.New("value reference is nil")
	}

	val := s.reflectValue.Interface()

	switch val.(type) {
	case int8:
		return int64(val.(int8)), nil
	case int16:
		return int64(val.(int16)), nil
	case int32:
		return int64(val.(int32)), nil
	case int:
		return int64(val.(int)), nil
	default:
		return val.(int64), nil
	}
}

func (s *signedIntegerType) SetValue(v int64) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	if s.Overflow(v) {
		return errors.New("value is too large")
	}

	switch s.reflectType.Name() {
	case "int8":
		s.reflectValue.Set(reflect.ValueOf(int8(v)))
	case "int16":
		s.reflectValue.Set(reflect.ValueOf(int16(v)))
	case "int32":
		s.reflectValue.Set(reflect.ValueOf(int32(v)))
	case "int":
		s.reflectValue.Set(reflect.ValueOf(int(v)))
	default:
		s.reflectValue.Set(reflect.ValueOf(int64(v)))
	}

	return nil
}

func (s *signedIntegerType) Overflow(v int64) bool {
	overflow := false

	if BitSize8 == s.bitSize && (math.MinInt8 > v || math.MaxInt8 < v) {
		overflow = true
	} else if BitSize16 == s.bitSize && (math.MinInt16 > v || math.MaxInt16 < v) {
		overflow = true
	} else if BitSize32 == s.bitSize && (math.MinInt32 > v || math.MaxInt32 < v) {
		overflow = true
	}

	return overflow
}

func (s *signedIntegerType) Instantiate() Value {
	return &value{
		reflect.New(s.reflectType),
	}
}

type UnsignedInteger interface {
	Type
	Instantiable
	BitSize() BitSize
	CanSet() bool
	Value() (uint64, error)
	SetValue(v uint64) error
	Overflow(v uint64) bool
}

type unsignedIntegerType struct {
	bitSize BitSize

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (u *unsignedIntegerType) Name() string {
	return u.reflectType.Name()
}

func (u *unsignedIntegerType) PackageName() string {
	return ""
}

func (u *unsignedIntegerType) PackagePath() string {
	return ""
}

func (u *unsignedIntegerType) HasValue() bool {
	return u.reflectValue != nil
}

func (u *unsignedIntegerType) Parent() Type {
	return u.parent
}

func (u *unsignedIntegerType) ReflectType() reflect.Type {
	return u.reflectType
}

func (u *unsignedIntegerType) ReflectValue() *reflect.Value {
	return u.reflectValue
}

func (u *unsignedIntegerType) BitSize() BitSize {
	return u.bitSize
}

func (u *unsignedIntegerType) CanSet() bool {
	if u.reflectValue == nil {
		return false
	}

	return u.reflectValue.CanSet()
}

func (u *unsignedIntegerType) Value() (uint64, error) {
	if u.reflectValue == nil {
		return 0, errors.New("value reference is nil")
	}

	val := u.reflectValue.Interface()

	switch val.(type) {
	case uint8:
		return uint64(val.(uint8)), nil
	case uint16:
		return uint64(val.(uint16)), nil
	case uint32:
		return uint64(val.(uint32)), nil
	case uint:
		return uint64(val.(uint)), nil
	default:
		return val.(uint64), nil
	}
}

func (u *unsignedIntegerType) SetValue(v uint64) error {
	if !u.CanSet() {
		return errors.New("value cannot be set")
	}

	if u.Overflow(v) {
		return errors.New("value is too large")
	}

	switch u.reflectType.Name() {
	case "uint8":
		u.reflectValue.Set(reflect.ValueOf(uint8(v)))
	case "uint16":
		u.reflectValue.Set(reflect.ValueOf(uint16(v)))
	case "uint32":
		u.reflectValue.Set(reflect.ValueOf(uint32(v)))
	case "uint":
		u.reflectValue.Set(reflect.ValueOf(uint(v)))
	default:
		u.reflectValue.Set(reflect.ValueOf(v))
	}

	return nil
}

func (u *unsignedIntegerType) Overflow(v uint64) bool {
	overflow := false

	if BitSize8 == u.bitSize && math.MaxUint8 < v {
		overflow = true
	} else if BitSize16 == u.bitSize && math.MaxUint16 < v {
		overflow = true
	} else if BitSize32 == u.bitSize && math.MaxUint32 < v {
		overflow = true
	}

	return overflow
}

func (u *unsignedIntegerType) Instantiate() Value {
	return &value{
		reflect.New(u.reflectType),
	}
}

type Float interface {
	Type
	Instantiable
	BitSize() BitSize
	CanSet() bool
	Value() (float64, error)
	SetValue(v float64) error
	Overflow(v float64) bool
}

type floatType struct {
	bitSize BitSize

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (f *floatType) Name() string {
	return f.reflectType.Name()
}

func (f *floatType) PackageName() string {
	return ""
}

func (f *floatType) PackagePath() string {
	return ""
}

func (f *floatType) HasValue() bool {
	return f.reflectValue != nil
}

func (f *floatType) Parent() Type {
	return f.parent
}

func (f *floatType) ReflectType() reflect.Type {
	return f.reflectType
}

func (f *floatType) ReflectValue() *reflect.Value {
	return f.reflectValue
}

func (f *floatType) BitSize() BitSize {
	return f.bitSize
}

func (f *floatType) Value() (float64, error) {
	if f.reflectValue == nil {
		return 0, errors.New("value reference is nil")
	}

	val := f.reflectValue.Interface()

	switch val.(type) {
	case float32:
		return float64(val.(float32)), nil
	default:
		return val.(float64), nil
	}
}

func (f *floatType) CanSet() bool {
	if f.reflectValue == nil {
		return false
	}

	return f.reflectValue.CanSet()
}

func (f *floatType) SetValue(v float64) error {
	if !f.CanSet() {
		return errors.New("value cannot be set")
	}

	if f.Overflow(v) {
		return errors.New("value is too large")
	}

	switch f.reflectType.Name() {
	case "float32":
		f.reflectValue.Set(reflect.ValueOf(float32(v)))
	default:
		f.reflectValue.Set(reflect.ValueOf(v))
	}

	return nil
}

func (f *floatType) Overflow(v float64) bool {
	overflow := false
	if BitSize32 == f.bitSize && math.MaxFloat32 < v {
		overflow = true
	} else if BitSize64 == f.bitSize && math.MaxFloat64 < v {
		overflow = true
	}
	return overflow
}

func (f *floatType) Instantiate() Value {
	return &value{
		reflect.New(f.reflectType),
	}
}

type Complex interface {
	Type
	Instantiable
	BitSize() BitSize
	CanSet() bool
	Value() (complex128, error)
	SetValue(v complex128) error
	ImaginaryData() (float64, error)
	RealData() (float64, error)
	SetImaginaryData(val float64) error
	SetRealData(val float64) error
}

type complexType struct {
	bitSize BitSize

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (c *complexType) Name() string {
	return c.reflectType.Name()
}

func (c *complexType) PackageName() string {
	return ""
}

func (c *complexType) PackagePath() string {
	return ""
}

func (c *complexType) HasValue() bool {
	return c.reflectValue != nil
}

func (c *complexType) Parent() Type {
	return c.parent
}

func (c *complexType) ReflectType() reflect.Type {
	return c.reflectType
}

func (c *complexType) ReflectValue() *reflect.Value {
	return c.reflectValue
}

func (c *complexType) BitSize() BitSize {
	return c.bitSize
}

func (c *complexType) CanSet() bool {
	if c.reflectValue == nil {
		return false
	}

	return c.reflectValue.CanSet()
}

func (c *complexType) Value() (complex128, error) {
	if c.reflectValue == nil {
		return 0, errors.New("value reference is nil")
	}

	val := c.reflectValue.Interface()

	switch val.(type) {
	case complex64:
		return complex128(val.(complex64)), nil
	default:
		return val.(complex128), nil
	}
}

func (c *complexType) SetValue(v complex128) error {
	if !c.CanSet() {
		return errors.New("value cannot be set")
	}

	switch c.reflectType.Name() {
	case "complex64":
		c.reflectValue.Set(reflect.ValueOf(complex64(v)))
	default:
		c.reflectValue.Set(reflect.ValueOf(v))
	}

	return nil
}

func (c *complexType) ImaginaryData() (float64, error) {
	if c.reflectValue == nil {
		return 0, errors.New("value reference is nil")
	}

	val := c.reflectValue.Interface()

	switch val.(type) {
	case complex64:
		return float64(imag(val.(complex64))), nil
	default:
		return imag(val.(complex128)), nil
	}
}

func (c *complexType) RealData() (float64, error) {
	if c.reflectValue == nil {
		return 0, errors.New("value reference is nil")
	}

	val := c.reflectValue.Interface()

	switch val.(type) {
	case complex64:
		return float64(real(val.(complex64))), nil
	default:
		return real(val.(complex128)), nil
	}
}

func (c *complexType) SetImaginaryData(val float64) error {
	if !c.CanSet() {
		return errors.New("value cannot be set")
	}

	real, err := c.RealData()

	if err != nil {
		return err
	}

	c.reflectValue.SetComplex(complex(real, val))
	return nil
}

func (c *complexType) SetRealData(val float64) error {
	if !c.CanSet() {
		return errors.New("value cannot be set")
	}

	imaginary, err := c.ImaginaryData()

	if err != nil {
		return err
	}

	c.reflectValue.SetComplex(complex(val, imaginary))
	return nil
}

func (c *complexType) Instantiate() Value {
	return &value{
		reflect.New(c.reflectType),
	}
}
