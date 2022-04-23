package reflector

func IsPointer(typ Type) bool {
	_, ok := typ.(Pointer)
	return ok
}

func ToPointer(typ Type) (Pointer, bool) {
	if ptr, ok := typ.(Pointer); ok {
		return ptr, true
	}

	return nil, false
}

func IsBasic(typ Type) bool {
	return true
}

func ToBasic(typ Type) (Basic, bool) {
	return nil, false
}

func IsStruct(typ Type) bool {
	_, ok := typ.(Struct)
	return ok
}

func ToStruct(typ Type) (Struct, bool) {
	if structType, ok := typ.(Struct); ok {
		return structType, true
	}

	return nil, false
}

func IsInterface(typ Type) bool {
	return true
}

func ToInterface(typ Type) (Interface, bool) {
	return nil, false
}

func IsFunction(typ Type) bool {
	return true
}

func ToFunction(typ Type) (Function, bool) {
	return nil, false
}

func IsArray(typ Type) bool {
	return true
}

func ToArray(typ Type) (Array, bool) {
	return nil, false
}

func IsSlice(typ Type) bool {
	_, ok := typ.(Slice)
	return ok
}

func ToSlice(typ Type) (Slice, bool) {
	if sliceType, ok := typ.(Slice); ok {
		return sliceType, true
	}

	return nil, false
}

func IsMap(typ Type) bool {
	_, ok := typ.(Map)
	return ok
}

func ToMap(typ Type) (Map, bool) {
	if mapType, ok := typ.(Map); ok {
		return mapType, true
	}

	return nil, false
}

func IsString(typ Type) bool {
	_, ok := typ.(String)
	return ok
}

func ToString(typ Type) (String, bool) {
	if stringType, ok := typ.(String); ok {
		return stringType, true
	}

	return nil, false
}

func IsBoolean(typ Type) bool {
	_, ok := typ.(Boolean)
	return ok
}

func ToBoolean(typ Type) (Boolean, bool) {
	if stringType, ok := typ.(Boolean); ok {
		return stringType, true
	}

	return nil, false
}

func IsInstantiable(typ Type) bool {
	_, ok := typ.(Instantiable)
	return ok
}

func ToInstantiable(typ Type) (Instantiable, bool) {
	if instantiableType, ok := typ.(Instantiable); ok {
		return instantiableType, true
	}

	return nil, false
}
