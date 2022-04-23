package reflector

import "reflect"

type Type interface {
	HasReference() bool
	ReflectType() reflect.Type
	ReflectValue() *reflect.Value
}

func TypeOf[T any]() Type {
	i := *new(T)
	typ := reflect.TypeOf(i)
	return typeOf(typ, nil)
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
			reflectType: typ,
		}
	case reflect.Interface:
	case reflect.Map:
		return &mapType{
			key:          typeOf(typ.Key(), val),
			elem:         typeOf(typ.Elem(), val),
			reflectType:  typ,
			reflectValue: val,
		}
	case reflect.Array:
	case reflect.Slice:
		sliceType := &sliceType{
			reflectType:  typ,
			reflectValue: val,
		}
		return sliceType
	case reflect.Func:
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
