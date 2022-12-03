package reflector

import (
	"reflect"
)

type Type interface {
	Name() string
	PackageName() string
	PackagePath() string
	CanSet() bool
	HasValue() bool
	Value() (any, error)
	SetValue(val any) error
	Parent() Type
	ReflectType() reflect.Type
	ReflectValue() *reflect.Value
	Compare(another Type) bool
	IsInstantiable() bool
	Instantiate() (Value, error)
	CanConvert(another Type) bool
	Convert(another Type) (Value, error)
}

var builtinTypes = map[string]struct{}{
	"string":     {},
	"bool":       {},
	"int":        {},
	"int8":       {},
	"int16":      {},
	"int32":      {},
	"int64":      {},
	"uint":       {},
	"uint8":      {},
	"uint16":     {},
	"uint32":     {},
	"uint64":     {},
	"float32":    {},
	"float64":    {},
	"complex64":  {},
	"complex128": {},
}

func TypeOf[T any]() Type {
	typ := reflect.TypeOf((*T)(nil))
	return typeOf(typ, typ.Elem(), nil, nil)
}

func TypeOfAny[T any](obj T) Type {
	nilType := reflect.TypeOf((*T)(nil))
	typ := reflect.TypeOf(obj)

	if typ == nil {
		return nil
	}

	val := reflect.ValueOf(obj)
	return typeOf(nilType, typ, &val, nil)
}

func typeOf(nilType reflect.Type, typ reflect.Type, val *reflect.Value, parent Type) Type {
	switch typ.Kind() {
	case reflect.Ptr:
		ptr := &pointer{
			reflectType:  typ,
			reflectValue: val,
		}

		if parent != nil && val != nil {
			ptr.base = typeOf(nilType, typ.Elem(), val, ptr)
		} else if val != nil {
			elem := val.Elem()
			ptr.base = typeOf(nilType, typ.Elem(), &elem, ptr)
		} else {
			ptr.base = typeOf(nilType, typ.Elem(), nil, ptr)
		}

		baseTypeName := ptr.base.Name()
		pointerTypeName := ptr.reflectType.Name()

		if pointerTypeName != "" && pointerTypeName != baseTypeName {
			return getCustomType(typ, val, parent, ptr)
		}

		return ptr
	case reflect.Struct:
		structType := &structType{
			parent:       parent,
			nilType:      nilType,
			reflectType:  typ,
			reflectValue: val,
		}

		z := typ.PkgPath()
		y := typ.Name()
		x := typ.String()
		if x == "" || y == "" || z == "" {

		}
		return structType
	case reflect.Interface:
		interfaceType := &interfaceType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		return interfaceType
	case reflect.Map:
		mapType := &mapType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		mapType.key = typeOf(nil, typ.Key(), val, mapType)
		mapType.elem = typeOf(nil, typ.Elem(), val, mapType)

		if typ.Name() != "" {
			return getCustomType(typ, val, parent, mapType)
		}

		return mapType
	case reflect.Array:
		arrayType := &arrayType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		arrayType.elem = typeOf(nil, typ.Elem(), val, arrayType)

		if typ.Name() != "" {
			return getCustomType(typ, val, parent, arrayType)
		}

		return arrayType
	case reflect.Slice:
		sliceType := &sliceType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		sliceType.elem = typeOf(nil, typ.Elem(), val, sliceType)

		if typ.Name() != "" {
			return getCustomType(typ, val, parent, sliceType)
		}

		return sliceType
	case reflect.Func:
		funcType := &functionType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		if typ.Name() != "" {
			return getCustomType(typ, val, parent, funcType)
		}

		return funcType
	case reflect.Chan:
		chanType := &chanType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		chanType.elem = typeOf(nil, typ.Elem(), val, chanType)

		if typ.Name() != "" {
			return getCustomType(typ, val, parent, chanType)
		}

		return chanType
	case reflect.String:
		stringType := &stringType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		if _, exists := builtinTypes[typ.Name()]; !exists {
			return getCustomType(typ, val, parent, stringType)
		}

		return stringType
	case reflect.Bool:
		booleanType := &booleanType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		if _, exists := builtinTypes[typ.Name()]; !exists {
			return getCustomType(typ, val, parent, booleanType)
		}

		return booleanType
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		integerType := &signedIntegerType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		if _, exists := builtinTypes[typ.Name()]; !exists {
			return getCustomType(typ, val, parent, integerType)
		}

		return integerType
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		integerType := &unsignedIntegerType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}
		if _, exists := builtinTypes[typ.Name()]; !exists {
			return getCustomType(typ, val, parent, integerType)
		}

		return integerType
	case reflect.Float32, reflect.Float64:
		floatType := &floatType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		if _, exists := builtinTypes[typ.Name()]; !exists {
			return getCustomType(typ, val, parent, floatType)
		}

		return floatType
	case reflect.Complex64, reflect.Complex128:
		complexType := &complexType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}

		if _, exists := builtinTypes[typ.Name()]; !exists {
			return getCustomType(typ, val, parent, complexType)
		}

		return complexType
	default:
		return nil
	}
}

func getCustomType(typ reflect.Type, val *reflect.Value, parent Type, underlyingType Type) *customType {
	return &customType{
		parent:         parent,
		reflectType:    typ,
		reflectValue:   val,
		underlyingType: underlyingType,
	}
}
