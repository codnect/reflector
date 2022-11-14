package reflector

func IsPointer(typ Type) bool {
	_, ok := typ.(*pointer)
	return ok
}

func ToPointer(typ Type) Pointer {
	if ptr, ok := typ.(*pointer); ok {
		return ptr
	}

	return nil
}

func IsStruct(typ Type) bool {
	_, ok := typ.(*structType)
	return ok
}

func ToStruct(typ Type) Struct {
	if structTyp, ok := typ.(*structType); ok {
		return structTyp
	}

	return nil
}

func IsCustom(typ Type) bool {
	_, ok := typ.(*customType)
	return ok
}

func ToCustom(typ Type) Custom {
	if customTyp, ok := typ.(*customType); ok {
		return customTyp
	}

	return nil
}

func IsInterface(typ Type) bool {
	_, ok := typ.(*interfaceType)
	return ok
}

func ToInterface(typ Type) Interface {
	if interfaceTyp, ok := typ.(*interfaceType); ok {
		return interfaceTyp
	}

	return nil
}

func IsFunction(typ Type) bool {
	_, ok := typ.(*functionType)
	return ok
}

func ToFunction(typ Type) Function {
	if functionTyp, ok := typ.(*functionType); ok {
		return functionTyp
	}

	return nil
}

func IsArray(typ Type) bool {
	_, ok := typ.(*arrayType)
	return ok
}

func ToArray(typ Type) Array {
	if arrayTyp, ok := typ.(*arrayType); ok {
		return arrayTyp
	}

	return nil
}

func IsSlice(typ Type) bool {
	_, ok := typ.(*sliceType)
	return ok
}

func ToSlice(typ Type) Slice {
	if sliceTyp, ok := typ.(*sliceType); ok {
		return sliceTyp
	}

	return nil
}

func IsMap(typ Type) bool {
	_, ok := typ.(*mapType)
	return ok
}

func ToMap(typ Type) Map {
	if mapTyp, ok := typ.(*mapType); ok {
		return mapTyp
	}

	return nil
}

func IsString(typ Type) bool {
	_, ok := typ.(*stringType)
	return ok
}

func ToString(typ Type) String {
	if strTyp, ok := typ.(*stringType); ok {
		return strTyp
	}

	return nil
}

func IsBoolean(typ Type) bool {
	_, ok := typ.(*booleanType)
	return ok
}

func ToBoolean(typ Type) Boolean {
	if boolTyp, ok := typ.(*booleanType); ok {
		return boolTyp
	}

	return nil
}

func IsInteger(typ Type) bool {
	return IsUnsignedInteger(typ) || IsSignedInteger(typ)
}

func IsSignedInteger(typ Type) bool {
	_, ok := typ.(*signedIntegerType)
	return ok
}

func ToSignedInteger(typ Type) SignedInteger {
	if intType, ok := typ.(*signedIntegerType); ok {
		return intType
	}

	return nil
}

func IsUnsignedInteger(typ Type) bool {
	_, ok := typ.(*unsignedIntegerType)
	return ok
}

func ToUnsignedInteger(typ Type) UnsignedInteger {
	if intType, ok := typ.(*unsignedIntegerType); ok {
		return intType
	}

	return nil
}

func IsFloat(typ Type) bool {
	_, ok := typ.(*floatType)
	return ok
}

func ToFloat(typ Type) Float {
	if floatTyp, ok := typ.(*floatType); ok {
		return floatTyp
	}

	return nil
}

func IsComplex(typ Type) bool {
	_, ok := typ.(*complexType)
	return ok
}

func ToComplex(typ Type) Complex {
	if complexTyp, ok := typ.(*complexType); ok {
		return complexTyp
	}

	return nil
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

func ToChan(typ Type) Chan {
	if chanTyp, ok := typ.(*chanType); ok {
		return chanTyp
	}

	return nil
}
