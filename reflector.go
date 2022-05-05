package reflector

func IsPointer(typ Type) bool {
	_, ok := typ.(*pointer)
	return ok
}

func ToPointer(typ Type) (Pointer, bool) {
	if ptr, ok := typ.(*pointer); ok {
		return ptr, true
	}

	return nil, false
}

func IsStruct(typ Type) bool {
	_, ok := typ.(*structType)
	return ok
}

func ToStruct(typ Type) (Struct, bool) {
	if structType, ok := typ.(*structType); ok {
		return structType, true
	}

	return nil, false
}

func IsCustom(typ Type) bool {
	_, ok := typ.(*customType)
	return ok
}

func ToCustom(typ Type) (Custom, bool) {
	if customType, ok := typ.(*customType); ok {
		return customType, true
	}

	return nil, false
}

func IsInterface(typ Type) bool {
	_, ok := typ.(*interfaceType)
	return ok
}

func ToInterface(typ Type) (Interface, bool) {
	if interfaceType, ok := typ.(*interfaceType); ok {
		return interfaceType, true
	}

	return nil, false
}

func IsFunction(typ Type) bool {
	_, ok := typ.(*functionType)
	return ok
}

func ToFunction(typ Type) (Function, bool) {
	if functionType, ok := typ.(*functionType); ok {
		return functionType, true
	}

	return nil, false
}

func IsArray(typ Type) bool {
	_, ok := typ.(*arrayType)
	return ok
}

func ToArray(typ Type) (Array, bool) {
	if arrayType, ok := typ.(*arrayType); ok {
		return arrayType, true
	}

	return nil, false
}

func IsSlice(typ Type) bool {
	_, ok := typ.(*sliceType)
	return ok
}

func ToSlice(typ Type) (Slice, bool) {
	if sliceType, ok := typ.(*sliceType); ok {
		return sliceType, true
	}

	return nil, false
}

func IsMap(typ Type) bool {
	_, ok := typ.(*mapType)
	return ok
}

func ToMap(typ Type) (Map, bool) {
	if mapType, ok := typ.(*mapType); ok {
		return mapType, true
	}

	return nil, false
}

func IsString(typ Type) bool {
	_, ok := typ.(*stringType)
	return ok
}

func ToString(typ Type) (String, bool) {
	if stringType, ok := typ.(*stringType); ok {
		return stringType, true
	}

	return nil, false
}

func IsBoolean(typ Type) bool {
	_, ok := typ.(*booleanType)
	return ok
}

func ToBoolean(typ Type) (Boolean, bool) {
	if stringType, ok := typ.(*booleanType); ok {
		return stringType, true
	}

	return nil, false
}

func IsInteger(typ Type) bool {
	return IsUnsignedInteger(typ) || IsSignedInteger(typ)
}

func IsSignedInteger(typ Type) bool {
	_, ok := typ.(*signedIntegerType)
	return ok
}

func ToSignedInteger(typ Type) (SignedInteger, bool) {
	if signedIntegerType, ok := typ.(*signedIntegerType); ok {
		return signedIntegerType, true
	}

	return nil, false
}

func IsUnsignedInteger(typ Type) bool {
	_, ok := typ.(*unsignedIntegerType)
	return ok
}

func ToUnsignedInteger(typ Type) (UnsignedInteger, bool) {
	if unsignedIntegerType, ok := typ.(*unsignedIntegerType); ok {
		return unsignedIntegerType, true
	}

	return nil, false
}

func IsFloat(typ Type) bool {
	_, ok := typ.(*floatType)
	return ok
}

func ToFloat(typ Type) (Float, bool) {
	if floatType, ok := typ.(*floatType); ok {
		return floatType, true
	}

	return nil, false
}

func IsComplex(typ Type) bool {
	_, ok := typ.(*complexType)
	return ok
}

func ToComplex(typ Type) (Complex, bool) {
	if complexType, ok := typ.(*complexType); ok {
		return complexType, true
	}

	return nil, false
}

func IsNumber(typ Type) bool {
	return IsInteger(typ) || IsFloat(typ) || IsComplex(typ)
}

func IsBasic(typ Type) bool {
	return IsBoolean(typ) || IsString(typ) || IsNumber(typ)
}

func IsChan(typ Type) bool {
	_, ok := typ.(*chanType)
	return ok
}

func ToChan(typ Type) (Chan, bool) {
	if chanType, ok := typ.(*chanType); ok {
		return chanType, true
	}

	return nil, false
}
