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
	BitSize() BitSize
	IntegerValue() (int64, error)
	SetIntegerValue(v int64) error
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

func (s *signedIntegerType) CanSet() bool {
	if s.reflectValue == nil {
		return false
	}

	return s.reflectValue.CanSet()
}

func (s *signedIntegerType) HasValue() bool {
	return s.reflectValue != nil
}

func (s *signedIntegerType) Value() (any, error) {
	return s.IntegerValue()
}

func (s *signedIntegerType) SetValue(v any) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	switch typedVal := v.(type) {
	case int8:
		return s.SetIntegerValue(int64(typedVal))
	case int16:
		return s.SetIntegerValue(int64(typedVal))
	case int32:
		return s.SetIntegerValue(int64(typedVal))
	case int64:
		return s.SetIntegerValue(typedVal)
	case int:
		return s.SetIntegerValue(int64(typedVal))
	default:
		return errors.New("type is not valid")
	}
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

func (s *signedIntegerType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return s.reflectType == another.ReflectType()
}

func (s *signedIntegerType) IsInstantiable() bool {
	return true
}

func (s *signedIntegerType) Instantiate() (Value, error) {
	return &value{
		reflect.New(s.reflectType),
	}, nil
}

func (s *signedIntegerType) BitSize() BitSize {
	return s.bitSize
}

func (s *signedIntegerType) IntegerValue() (int64, error) {
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

func (s *signedIntegerType) SetIntegerValue(v int64) error {
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

type UnsignedInteger interface {
	Type
	BitSize() BitSize
	IntegerValue() (uint64, error)
	SetIntegerValue(v uint64) error
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

func (u *unsignedIntegerType) CanSet() bool {
	if u.reflectValue == nil {
		return false
	}

	return u.reflectValue.CanSet()
}

func (u *unsignedIntegerType) HasValue() bool {
	return u.reflectValue != nil
}

func (u *unsignedIntegerType) Value() (any, error) {
	return u.IntegerValue()
}

func (u *unsignedIntegerType) SetValue(v any) error {
	if !u.CanSet() {
		return errors.New("value cannot be set")
	}

	switch typedVal := v.(type) {
	case uint8:
		return u.SetIntegerValue(uint64(typedVal))
	case uint16:
		return u.SetIntegerValue(uint64(typedVal))
	case uint32:
		return u.SetIntegerValue(uint64(typedVal))
	case uint64:
		return u.SetIntegerValue(typedVal)
	case uint:
		return u.SetIntegerValue(uint64(typedVal))
	default:
		return errors.New("type is not valid")
	}
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

func (u *unsignedIntegerType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return u.reflectType == another.ReflectType()
}

func (u *unsignedIntegerType) IsInstantiable() bool {
	return true
}

func (u *unsignedIntegerType) Instantiate() (Value, error) {
	return &value{
		reflect.New(u.reflectType),
	}, nil
}

func (u *unsignedIntegerType) BitSize() BitSize {
	return u.bitSize
}

func (u *unsignedIntegerType) IntegerValue() (uint64, error) {
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

func (u *unsignedIntegerType) SetIntegerValue(v uint64) error {
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

type Float interface {
	Type
	BitSize() BitSize
	FloatValue() (float64, error)
	SetFloatValue(v float64) error
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

func (f *floatType) CanSet() bool {
	if f.reflectValue == nil {
		return false
	}

	return f.reflectValue.CanSet()
}

func (f *floatType) HasValue() bool {
	return f.reflectValue != nil
}

func (f *floatType) Value() (any, error) {
	return f.FloatValue()
}

func (f *floatType) SetValue(v any) error {
	if !f.CanSet() {
		return errors.New("value cannot be set")
	}

	switch typedVal := v.(type) {
	case float32:
		return f.SetFloatValue(float64(typedVal))
	case float64:
		return f.SetFloatValue(typedVal)
	default:
		return errors.New("type is not valid")
	}
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

func (f *floatType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return f.reflectType == another.ReflectType()
}

func (f *floatType) IsInstantiable() bool {
	return true
}

func (f *floatType) Instantiate() (Value, error) {
	return &value{
		reflect.New(f.reflectType),
	}, nil
}

func (f *floatType) BitSize() BitSize {
	return f.bitSize
}

func (f *floatType) FloatValue() (float64, error) {
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

func (f *floatType) SetFloatValue(v float64) error {
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

type Complex interface {
	Type
	BitSize() BitSize
	ComplexValue() (complex128, error)
	SetComplexValue(v complex128) error
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

func (c *complexType) CanSet() bool {
	if c.reflectValue == nil {
		return false
	}

	return c.reflectValue.CanSet()
}

func (c *complexType) HasValue() bool {
	return c.reflectValue != nil
}

func (c *complexType) Value() (any, error) {
	return c.ComplexValue()
}

func (c *complexType) SetValue(v any) error {
	if !c.CanSet() {
		return errors.New("value cannot be set")
	}

	switch typedVal := v.(type) {
	case complex64:
		return c.SetComplexValue(complex128(typedVal))
	case complex128:
		return c.SetComplexValue(typedVal)
	default:
		return errors.New("type is not valid")
	}
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

func (c *complexType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return c.reflectType == another.ReflectType()
}

func (c *complexType) IsInstantiable() bool {
	return true
}

func (c *complexType) Instantiate() (Value, error) {
	return &value{
		reflect.New(c.reflectType),
	}, nil
}

func (c *complexType) BitSize() BitSize {
	return c.bitSize
}

func (c *complexType) ComplexValue() (complex128, error) {
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

func (c *complexType) SetComplexValue(v complex128) error {
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
