package reflector

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeOfFunction(t *testing.T) {
	typ := TypeOf[func(param1 string, param2 []int, param3 ...any) (int, error)]()
	assert.True(t, IsFunction(typ))
	assert.Equal(t, "func(string,[]int,[]any) (int,error)", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	functionType := ToFunction(typ)

	assert.NotNil(t, functionType)

	assert.True(t, functionType.IsVariadic())
	//assert.False(t, functionType.HasReceiver())
	assert.False(t, functionType.IsExported())
	//receiverType, ok := functionType.Receiver()

	//assert.Nil(t, receiverType)
	//assert.False(t, ok)

	assert.Equal(t, 3, functionType.NumParameter())
	params := functionType.Parameters()
	assert.Len(t, params, 3)

	param := functionType.Parameters()[0]
	stringParamType := ToString(param)
	assert.Equal(t, "string", stringParamType.Name())

	param = functionType.Parameters()[1]
	sliceParamType := ToSlice(param)
	assert.Equal(t, "[]int", sliceParamType.Name())

	param = functionType.Parameters()[2]
	sliceParamType = ToSlice(param)
	assert.Equal(t, "[]any", sliceParamType.Name())

	assert.Equal(t, 2, functionType.NumResult())

	results := functionType.Results()
	assert.Len(t, results, 2)

	result := functionType.Results()[0]
	integerResultType := ToSignedInteger(result)
	assert.Equal(t, "int", integerResultType.Name())

	param = functionType.Results()[1]
	interfaceResultType := ToInterface(param)
	assert.Equal(t, "error", interfaceResultType.Name())

	outputs, err := functionType.Invoke("anyTestValue1", []int{2, 5}, "anyTestValue2", 6)
	assert.Nil(t, outputs)
	assert.NotNil(t, err)
}

func TestTypeOfFunctionPointer(t *testing.T) {
	ptrType := TypeOf[*func(param1 string, param2 []int, param3 ...any) (int, error)]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*func(string,[]int,[]any) (int,error)", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsFunction(typ))
	assert.Equal(t, "func(string,[]int,[]any) (int,error)", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	functionType := ToFunction(typ)

	assert.NotNil(t, functionType)

	assert.True(t, functionType.IsVariadic())
	//assert.False(t, functionType.HasReceiver())
	assert.False(t, functionType.IsExported())
	//receiverType, ok := functionType.Receiver()

	//assert.Nil(t, receiverType)
	//assert.False(t, ok)

	assert.Equal(t, 3, functionType.NumParameter())
	params := functionType.Parameters()
	assert.Len(t, params, 3)

	param := functionType.Parameters()[0]
	stringParamType := ToString(param)
	assert.Equal(t, "string", stringParamType.Name())

	param = functionType.Parameters()[1]
	sliceParamType := ToSlice(param)
	assert.Equal(t, "[]int", sliceParamType.Name())

	param = functionType.Parameters()[2]
	sliceParamType = ToSlice(param)
	assert.Equal(t, "[]any", sliceParamType.Name())

	assert.Equal(t, 2, functionType.NumResult())

	results := functionType.Results()
	assert.Len(t, results, 2)

	result := functionType.Results()[0]
	integerResultType := ToSignedInteger(result)
	assert.Equal(t, "int", integerResultType.Name())

	param = functionType.Results()[1]
	interfaceResultType := ToInterface(param)
	assert.Equal(t, "error", interfaceResultType.Name())

	outputs, err := functionType.Invoke("anyTestValue1", []int{2, 5}, "anyTestValue2", 6)
	assert.Nil(t, outputs)
	assert.NotNil(t, err)
}

func TestTypeOfFunctionObject(t *testing.T) {
	var val func(param1 string, param2 []int, param3 ...any) (int, error)

	typ := TypeOfAny(val)
	assert.True(t, IsFunction(typ))
	assert.Equal(t, "func(string,[]int,[]any) (int,error)", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	functionType := ToFunction(typ)

	assert.NotNil(t, functionType)

	assert.True(t, functionType.IsVariadic())
	//assert.False(t, functionType.HasReceiver())
	assert.False(t, functionType.IsExported())
	//receiverType, ok := functionType.Receiver()

	//assert.Nil(t, receiverType)
	//assert.False(t, ok)

	assert.Equal(t, 3, functionType.NumParameter())
	params := functionType.Parameters()
	assert.Len(t, params, 3)

	param := functionType.Parameters()[0]
	stringParamType := ToString(param)
	assert.Equal(t, "string", stringParamType.Name())

	param = functionType.Parameters()[1]
	sliceParamType := ToSlice(param)
	assert.Equal(t, "[]int", sliceParamType.Name())

	param = functionType.Parameters()[2]
	sliceParamType = ToSlice(param)
	assert.Equal(t, "[]any", sliceParamType.Name())

	assert.Equal(t, 2, functionType.NumResult())

	results := functionType.Results()
	assert.Len(t, results, 2)

	result := functionType.Results()[0]
	integerResultType := ToSignedInteger(result)
	assert.Equal(t, "int", integerResultType.Name())

	param = functionType.Results()[1]
	interfaceResultType := ToInterface(param)
	assert.Equal(t, "error", interfaceResultType.Name())

	outputs, err := functionType.Invoke("anyTestValue1", []int{2, 5}, "anyTestValue2", 6)
	assert.Nil(t, outputs)
	assert.NotNil(t, err)
}

func TestTypeOfFunctionObjectPointer(t *testing.T) {
	var val func(param1 string, param2 []int, param3 ...any) (int, error)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*func(string,[]int,[]any) (int,error)", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsFunction(typ))
	assert.Equal(t, "func(string,[]int,[]any) (int,error)", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	functionType := ToFunction(typ)

	assert.NotNil(t, functionType)

	assert.True(t, functionType.IsVariadic())
	//assert.False(t, functionType.HasReceiver())
	assert.False(t, functionType.IsExported())
	//receiverType, ok := functionType.Receiver()

	//assert.Nil(t, receiverType)
	//assert.False(t, ok)

	assert.Equal(t, 3, functionType.NumParameter())
	params := functionType.Parameters()
	assert.Len(t, params, 3)

	param := functionType.Parameters()[0]
	stringParamType := ToString(param)
	assert.Equal(t, "string", stringParamType.Name())

	param = functionType.Parameters()[1]
	sliceParamType := ToSlice(param)
	assert.Equal(t, "[]int", sliceParamType.Name())

	param = functionType.Parameters()[2]
	sliceParamType = ToSlice(param)
	assert.Equal(t, "[]any", sliceParamType.Name())

	assert.Equal(t, 2, functionType.NumResult())

	results := functionType.Results()
	assert.Len(t, results, 2)

	result := functionType.Results()[0]
	integerResultType := ToSignedInteger(result)
	assert.Equal(t, "int", integerResultType.Name())

	param = functionType.Results()[1]
	interfaceResultType := ToInterface(param)
	assert.Equal(t, "error", interfaceResultType.Name())

	outputs, err := functionType.Invoke("anyTestValue1", []int{2, 5}, "anyTestValue2", 6)
	assert.Nil(t, outputs)
	assert.NotNil(t, err)
}

func Function1(param1 *string, param2 []int, param3 string, param4 ...any) (int, error) {
	return 25, errors.New("Function1")
}

func TestTypeOfTestFunction(t *testing.T) {
	var val func(param1 *string, param2 []int, param3 string, param4 ...any) (int, error)
	val = Function1

	typ := TypeOfAny(val)
	assert.True(t, IsFunction(typ))
	assert.Equal(t, "Function1", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	functionType := ToFunction(typ)

	assert.NotNil(t, functionType)

	assert.True(t, functionType.IsVariadic())
	//assert.False(t, functionType.HasReceiver())
	assert.False(t, functionType.IsExported())
	//receiverType, ok := functionType.Receiver()

	//assert.Nil(t, receiverType)
	//assert.False(t, ok)

	assert.Equal(t, 4, functionType.NumParameter())
	params := functionType.Parameters()
	assert.Len(t, params, 4)

	param := functionType.Parameters()[0]
	assert.True(t, IsPointer(param))
	ptr := ToPointer(param)

	assert.NotNil(t, ptr)

	stringParamType := ToString(ptr.Elem())
	assert.Equal(t, "string", stringParamType.Name())

	param = functionType.Parameters()[1]
	sliceParamType := ToSlice(param)
	assert.Equal(t, "[]int", sliceParamType.Name())

	param = functionType.Parameters()[2]
	stringParamType = ToString(param)
	assert.Equal(t, "string", stringParamType.Name())

	param = functionType.Parameters()[3]
	sliceParamType = ToSlice(param)
	assert.Equal(t, "[]any", sliceParamType.Name())

	assert.Equal(t, 2, functionType.NumResult())

	results := functionType.Results()
	assert.Len(t, results, 2)

	result := functionType.Results()[0]
	integerResultType := ToSignedInteger(result)
	assert.Equal(t, "int", integerResultType.Name())

	param = functionType.Results()[1]
	interfaceResultType := ToInterface(param)
	assert.Equal(t, "error", interfaceResultType.Name())

	outputs, err := functionType.Invoke(new(string), []int{2, 5}, nil, nil, 6, 8, 10)
	assert.NotNil(t, outputs)
	assert.Nil(t, err)

	assert.Equal(t, 25, outputs[0])
	assert.Equal(t, "Function1", outputs[1].(error).Error())
}
