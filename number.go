package reflector

import (
	"errors"
	"math"
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

func toBitSize(bits int) BitSize {
	switch bits {
	case 8:
		return BitSize8
	case 16:
		return BitSize16
	case 32:
		return BitSize32
	case 64:
		return BitSize64
	default:
		return BitSize128
	}
}

type SignedInteger interface {
	Type
	BitSize() BitSize
	IntegerValue() (int64, error)
	SetIntegerValue(v int64) error
	Overflow(v int64) bool
}

type signedIntegerType struct {
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
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32, float64:
		convertedValue := reflect.ValueOf(typedVal).Convert(s.reflectType)
		s.reflectValue.Set(convertedValue)
		return nil
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

func (s *signedIntegerType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if s.reflectValue == nil {
		return s.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return s.reflectValue.CanConvert(typ.ReflectType())
}

func (s *signedIntegerType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if s.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !s.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := s.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (s *signedIntegerType) BitSize() BitSize {
	return toBitSize(s.reflectType.Bits())
}

func (s *signedIntegerType) IntegerValue() (int64, error) {
	if s.reflectValue == nil {
		return -1, errors.New("value reference is nil")
	}

	return s.reflectValue.Convert(TypeOf[int64]().ReflectType()).Interface().(int64), nil
}

func (s *signedIntegerType) SetIntegerValue(v int64) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	s.reflectValue.Set(reflect.ValueOf(v).Convert(s.reflectType))
	return nil
}

func (s *signedIntegerType) Overflow(v int64) bool {
	overflow := false

	bitSize := s.BitSize()

	if BitSize8 == bitSize && (math.MinInt8 > v || math.MaxInt8 < v) {
		overflow = true
	} else if BitSize16 == bitSize && (math.MinInt16 > v || math.MaxInt16 < v) {
		overflow = true
	} else if BitSize32 == bitSize && (math.MinInt32 > v || math.MaxInt32 < v) {
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
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32, float64:
		convertedValue := reflect.ValueOf(typedVal).Convert(u.reflectType)
		u.reflectValue.Set(convertedValue)
		return nil
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

func (u *unsignedIntegerType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if u.reflectValue == nil {
		return u.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return u.reflectValue.CanConvert(typ.ReflectType())
}

func (u *unsignedIntegerType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if u.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !u.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := u.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (u *unsignedIntegerType) BitSize() BitSize {
	return toBitSize(u.reflectType.Bits())
}

func (u *unsignedIntegerType) IntegerValue() (uint64, error) {
	if u.reflectValue == nil {
		return 0, errors.New("value reference is nil")
	}

	return u.reflectValue.Convert(TypeOf[uint64]().ReflectType()).Interface().(uint64), nil
}

func (u *unsignedIntegerType) SetIntegerValue(v uint64) error {
	if !u.CanSet() {
		return errors.New("value cannot be set")
	}

	u.reflectValue.Set(reflect.ValueOf(v).Convert(u.reflectType))
	return nil
}

func (u *unsignedIntegerType) Overflow(v uint64) bool {
	overflow := false

	bitSize := u.BitSize()
	if BitSize8 == bitSize && math.MaxUint8 < v {
		overflow = true
	} else if BitSize16 == bitSize && math.MaxUint16 < v {
		overflow = true
	} else if BitSize32 == bitSize && math.MaxUint32 < v {
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
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32, float64:
		convertedValue := reflect.ValueOf(typedVal).Convert(f.reflectType)
		f.reflectValue.Set(convertedValue)
		return nil
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

func (f *floatType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if f.reflectValue == nil {
		return f.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return f.reflectValue.CanConvert(typ.ReflectType())
}

func (f *floatType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if f.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !f.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := f.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (f *floatType) BitSize() BitSize {
	return toBitSize(f.reflectType.Bits())
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

	f.reflectValue.Set(reflect.ValueOf(v).Convert(f.reflectType))
	return nil
}

func (f *floatType) Overflow(v float64) bool {
	overflow := false
	bitSize := f.BitSize()
	if BitSize32 == bitSize && math.MaxFloat32 < v {
		overflow = true
	} else if BitSize64 == bitSize && math.MaxFloat64 < v {
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

func (c *complexType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if c.reflectValue == nil {
		return c.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return c.reflectValue.CanConvert(typ.ReflectType())
}

func (c *complexType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if c.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !c.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := c.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (c *complexType) BitSize() BitSize {
	return toBitSize(c.reflectType.Bits())
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
