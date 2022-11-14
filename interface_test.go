package reflector

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
	_ "testing"
)

type TestInterface interface {
	Find(name string) (string, error)
	print() string
}

type TestGenericInterface[T any] interface {
	Find(name bool) (T, error)
	print() string
}

type TestStruct4 struct {
}

func (x TestStruct4) Find(name string) (string, error) {
	return "", nil
}

func (x TestStruct4) print() string {
	return ""
}

func TestTypeOfInterface(t *testing.T) {
	typ := TypeOf[TestInterface]()
	assert.True(t, IsInterface(typ))
	assert.Equal(t, "TestInterface", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	iface := ToInterface(typ)

	assert.NotNil(t, iface)

	assert.Equal(t, 2, iface.NumMethod())

	methods := iface.Methods()
	assert.Len(t, methods, 2)

	method := iface.Methods()[0]
	assert.NotNil(t, method)

	assert.Equal(t, "Find", method.Name())
	assert.True(t, method.IsExported())
	assert.Equal(t, 1, method.NumParameter())
	assert.Equal(t, 2, method.NumResult())
	//assert.False(t, method1.HasReceiver())

	method1Parameters := method.Parameters()
	assert.Len(t, method1Parameters, 1)

	assert.NotNil(t, method1Parameters[0])
	assert.Equal(t, "string", method1Parameters[0].Name())
	assert.Equal(t, "", method1Parameters[0].PackageName())

	method1Results := method.Results()
	assert.Len(t, method1Results, 2)

	assert.NotNil(t, method1Results[0])
	assert.Equal(t, "string", method1Results[0].Name())
	assert.Equal(t, "", method1Results[0].PackageName())

	assert.NotNil(t, method1Results[1])
	assert.Equal(t, "error", method1Results[1].Name())
	assert.Equal(t, "", method1Results[1].PackageName())

	method = iface.Methods()[1]
	assert.NotNil(t, method)

	assert.Equal(t, "print", method.Name())
	assert.False(t, method.IsExported())
	assert.Equal(t, 0, method.NumParameter())
	assert.Equal(t, 1, method.NumResult())
	//assert.False(t, method2.HasReceiver())

	method2Parameters := method.Parameters()
	assert.Len(t, method2Parameters, 0)

	method2Results := method.Results()
	assert.Len(t, method2Results, 1)

	assert.NotNil(t, method2Results[0])
	assert.Equal(t, "string", method2Results[0].Name())
	assert.Equal(t, "", method2Results[0].PackageName())
}

func TestTypeOfGenericInterface(t *testing.T) {
	typ := TypeOf[TestGenericInterface[int]]()
	assert.True(t, IsInterface(typ))
	assert.Equal(t, "TestGenericInterface[int]", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	iface := ToInterface(typ)

	assert.NotNil(t, iface)

	assert.Equal(t, 2, iface.NumMethod())

	methods := iface.Methods()
	assert.Len(t, methods, 2)

	method1 := iface.Methods()[0]
	assert.NotNil(t, method1)

	assert.Equal(t, "Find", method1.Name())
	assert.True(t, method1.IsExported())
	assert.Equal(t, 1, method1.NumParameter())
	assert.Equal(t, 2, method1.NumResult())
	//assert.False(t, method1.HasReceiver())

	method1Parameters := method1.Parameters()
	assert.Len(t, method1Parameters, 1)

	assert.NotNil(t, method1Parameters[0])
	assert.Equal(t, "bool", method1Parameters[0].Name())
	assert.Equal(t, "", method1Parameters[0].PackageName())

	method1Results := method1.Results()
	assert.Len(t, method1Results, 2)

	assert.NotNil(t, method1Results[0])
	assert.Equal(t, "int", method1Results[0].Name())
	assert.Equal(t, "", method1Results[0].PackageName())

	assert.NotNil(t, method1Results[1])
	assert.Equal(t, "error", method1Results[1].Name())
	assert.Equal(t, "", method1Results[1].PackageName())

	method2 := iface.Methods()[1]
	assert.NotNil(t, method2)

	assert.Equal(t, "print", method2.Name())
	assert.False(t, method2.IsExported())
	assert.Equal(t, 0, method2.NumParameter())
	assert.Equal(t, 1, method2.NumResult())
	//assert.False(t, method2.HasReceiver())

	method2Parameters := method2.Parameters()
	assert.Len(t, method2Parameters, 0)

	method2Results := method2.Results()
	assert.Len(t, method2Results, 1)

	assert.NotNil(t, method2Results[0])
	assert.Equal(t, "string", method2Results[0].Name())
	assert.Equal(t, "", method2Results[0].PackageName())
}

func TestTypeOfInterfaceObject(t *testing.T) {
	var interfaceObject TestInterface = TestStruct4{}
	typ := TypeOfAny(interfaceObject)
	assert.True(t, IsInterface(typ))
	assert.Equal(t, "TestInterface", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	iface := ToInterface(typ)

	assert.NotNil(t, iface)

	assert.Equal(t, 2, iface.NumMethod())

	methods := iface.Methods()
	assert.Len(t, methods, 2)

	method := iface.Methods()[0]
	assert.NotNil(t, method)

	assert.Equal(t, "Find", method.Name())
	assert.True(t, method.IsExported())
	assert.Equal(t, 1, method.NumParameter())
	assert.Equal(t, 2, method.NumResult())
	//assert.False(t, method1.HasReceiver())

	method1Parameters := method.Parameters()
	assert.Len(t, method1Parameters, 1)

	method.Invoke("anyValue")

	assert.NotNil(t, method1Parameters[0])
	assert.Equal(t, "string", method1Parameters[0].Name())
	assert.Equal(t, "", method1Parameters[0].PackageName())

	method1Results := method.Results()
	assert.Len(t, method1Results, 2)

	assert.NotNil(t, method1Results[0])
	assert.Equal(t, "string", method1Results[0].Name())
	assert.Equal(t, "", method1Results[0].PackageName())

	assert.NotNil(t, method1Results[1])
	assert.Equal(t, "error", method1Results[1].Name())
	assert.Equal(t, "", method1Results[1].PackageName())

	method = iface.Methods()[1]
	assert.NotNil(t, method)

	assert.Equal(t, "print", method.Name())
	assert.False(t, method.IsExported())
	assert.Equal(t, 0, method.NumParameter())
	assert.Equal(t, 1, method.NumResult())
	//assert.False(t, method2.HasReceiver())

	method2Parameters := method.Parameters()
	assert.Len(t, method2Parameters, 0)

	method2Results := method.Results()
	assert.Len(t, method2Results, 1)

	assert.NotNil(t, method2Results[0])
	assert.Equal(t, "string", method2Results[0].Name())
	assert.Equal(t, "", method2Results[0].PackageName())
}
