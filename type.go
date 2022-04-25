package reflector

import "reflect"

type Type interface {
	Name() string
	PackageName() string
	HasValue() bool
	Parent() Type
	ReflectType() reflect.Type
	ReflectValue() *reflect.Value
}

func TypeOf[T any]() Type {
	iface := (*T)(nil)
	typ := reflect.TypeOf(iface)
	return typeOf(typ, nil, nil).(Pointer).Elem()
}

func TypeOfAny(obj any) Type {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	return typeOf(typ, &val, nil)
}

func typeOf(typ reflect.Type, val *reflect.Value, parent Type) Type {

	switch typ.Kind() {
	case reflect.Ptr:
		ptr := &pointer{
			reflectType:  typ,
			reflectValue: val,
		}

		if val != nil {
			elem := val.Elem()
			ptr.base = typeOf(typ.Elem(), &elem, ptr)
		} else {
			ptr.base = typeOf(typ.Elem(), nil, ptr)
		}

		return ptr
	case reflect.Struct:
		return &structType{
			parent:       parent,
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
			key:          typeOf(typ.Key(), val, nil),
			elem:         typeOf(typ.Elem(), val, nil),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Array:
		sliceType := &arrayType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,

			elem: typeOf(typ.Elem(), val, nil),
		}
		return sliceType
	case reflect.Slice:
		sliceType := &sliceType{
			parent:       parent,
			reflectType:  typ,
			reflectValue: val,
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
	}

	return nil
}
