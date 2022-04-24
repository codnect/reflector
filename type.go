package reflector

import "reflect"

type Type interface {
	Name() string
	PackageName() string
	HasValue() bool
	ReflectType() reflect.Type
	ReflectValue() *reflect.Value
}

func TypeOf[T any]() Type {
	iface := (*T)(nil)
	typ := reflect.TypeOf(iface)
	return typeOf(typ, nil).(Pointer).Elem()
}

func TypeOfAny(obj any) Type {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	return typeOf(typ, &val)
}

func typeOf(typ reflect.Type, val *reflect.Value) Type {

	switch typ.Kind() {
	case reflect.Ptr:
		ptr := &pointer{
			reflectType:  typ,
			reflectValue: val,
		}

		if val != nil {
			elem := val.Elem()
			ptr.base = typeOf(typ.Elem(), &elem)
		} else {
			ptr.base = typeOf(typ.Elem(), nil)
		}

		return ptr
	case reflect.Struct:
		return &structType{
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Interface:
		return &interfaceType{
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Map:
		return &mapType{
			key:          typeOf(typ.Key(), val),
			elem:         typeOf(typ.Elem(), val),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Array:
		sliceType := &arrayType{
			reflectType:  typ,
			reflectValue: val,

			elem: typeOf(typ.Elem(), val),
		}
		return sliceType
	case reflect.Slice:
		sliceType := &sliceType{
			reflectType:  typ,
			reflectValue: val,
		}
		return sliceType
	case reflect.Func:
		return &functionType{
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Chan:
	case reflect.String:
		return &stringType{
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Bool:
		return &booleanType{
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &signedInteger{
			bitSize:      bitSize(typ.Kind()),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &unsignedInteger{
			bitSize:      bitSize(typ.Kind()),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Float32, reflect.Float64:
		return &float{
			bitSize:      bitSize(typ.Kind()),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Complex64, reflect.Complex128:
		return &complexType{
			bitSize:      bitSize(typ.Kind()),
			reflectType:  typ,
			reflectValue: val,
		}
	}

	return nil
}

/*
type Type struct {
	typ any
}


func (t *Type) FullName() string {
	return ""
}


func TypeOfConstructor(obj any) *Type {
	return &Type{}
}

func TypeOfAnyObject(obj any) *Type {
	return &Type{}
}
*/
