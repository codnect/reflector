package reflector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test interface {
	Print()
}

type Wrapper struct {
	Name string
}

type Dessert struct {
	Wrapper Wrapper
}

func (d Dessert) Print() {

}

func (d Dessert) Eat() string {
	return ""
}

func (d Dessert) Buy() {

}

func TestTypeOfStruct(t *testing.T) {

	typ := TypeOf[Dessert]()
	assert.True(t, IsStruct(typ))
	assert.Equal(t, "Dessert", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	structType, isStruct := ToStruct(typ)

	assert.NotNil(t, structType)
	assert.True(t, isStruct)

	//	assert.Equal(t, 2, structType.NumMethod())

	y := TypeOf[Test]()
	i, _ := ToInterface(y)
	b := structType.Implements(i)
	if b {

	}
}

func TestTypeOfStructPointer(t *testing.T) {
	ptrType := TypeOf[*Dessert]()

	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*Dessert", ptr.Name())
	assert.Equal(t, "reflector", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsStruct(typ))
	assert.Equal(t, "Dessert", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	structType, isStruct := ToStruct(typ)

	assert.NotNil(t, structType)
	assert.True(t, isStruct)

	assert.Equal(t, 3, structType.NumMethod())
	structType.Methods()

	y := TypeOf[Test]()
	i, _ := ToInterface(y)
	b := structType.Implements(i)
	if b {

	}

}

func TestTypeOfStructObject(t *testing.T) {
	dessert := Dessert{
		Wrapper: Wrapper{
			Name: "burak",
		},
	}

	ptrType := TypeOfAny(&dessert)

	ptr, _ := ToPointer(ptrType)
	typ := ptr.Elem()

	assert.True(t, IsStruct(typ))
	assert.Equal(t, "Dessert", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	structType, isStruct := ToStruct(typ)

	assert.NotNil(t, structType)
	assert.True(t, isStruct)
	assert.True(t, structType.CanSet())

	e := structType.Fields()[0]
	s, _ := ToStruct(e.Type())

	ft := s.Fields()[0].Type()
	str, _ := ToString(ft)
	str.SetValue("hello")

	v := s.Fields()[0].CanSet()

	if v {

	}

	//	assert.Equal(t, 2, structType.NumMethod())

	y := TypeOf[Test]()
	i, _ := ToInterface(y)
	b := structType.Implements(i)
	if b {

	}
}
