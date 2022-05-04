package reflector

import (
	"reflect"
)

type Type interface {
	Name() string
	PackageName() string
	PackagePath() string
	HasValue() bool
	Parent() Type
	ReflectType() reflect.Type
	ReflectValue() *reflect.Value
	Compare(another Type) bool
}

func TypeOf[T any]() Type {
	typ := reflect.TypeOf((*T)(nil))
	return typeOf(typ, typ.Elem(), nil, nil)
}

func TypeOfAny[T any](obj T) Type {
	nilType := reflect.TypeOf((*T)(nil))

	typ := reflect.TypeOf(obj)
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

		if val != nil {
			elem := val.Elem()
			ptr.base = typeOf(nilType, typ.Elem(), &elem, ptr)
		} else {
			ptr.base = typeOf(nilType, typ.Elem(), nil, ptr)
		}

		return ptr
	case reflect.Struct:
		return &structType{
			parent:       parent,
			nilType:      nilType,
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Interface:
		return &interfaceType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Map:
		return &mapType{
			parent:       parent,
			key:          typeOf(nil, typ.Key(), val, nil),
			elem:         typeOf(nil, typ.Elem(), val, nil),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Array:
		arrayType := &arrayType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,

			elem: typeOf(nil, typ.Elem(), val, nil),
		}
		return arrayType
	case reflect.Slice:
		sliceType := &sliceType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,

			elem: typeOf(nil, typ.Elem(), val, nil),
		}
		return sliceType
	case reflect.Func:
		return &functionType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Chan:
		return &chanType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
			elem:         typeOf(nil, typ.Elem(), val, nil),
		}
	case reflect.String:
		return &stringType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Bool:
		return &booleanType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &signedIntegerType{
			parent:       parent,
			bitSize:      bitSize(typ.Kind()),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &unsignedIntegerType{
			parent:       parent,
			bitSize:      bitSize(typ.Kind()),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Float32, reflect.Float64:
		return &floatType{
			parent:       parent,
			bitSize:      bitSize(typ.Kind()),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Complex64, reflect.Complex128:
		return &complexType{
			parent:       parent,
			bitSize:      bitSize(typ.Kind()),
			reflectType:  typ,
			reflectValue: val,
		}
	default:
		return &customType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
		}
	}
}
