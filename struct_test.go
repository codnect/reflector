package reflector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestInterface1 interface {
	Method1() int
	Method2(val string)
}

type TestInterface2 interface {
	Method3() error
}

type TestInterface3 interface {
	Method4()
}

type TestInterface4 interface {
	Method5()
	Method6()
}

type TestInterface5 interface {
	Method7()
}

type TestEmbeddedStruct struct {
	ExportedAndEmbeddedField   byte
	unexportedAndEmbeddedField int16
}

type TestStruct1 struct {
	TestInterface5       `json:"TestInterface5"`
	TestEmbeddedStruct   `json:"TestEmbeddedStruct"`
	ExportedStructField  TestStruct2 `json:"ExportedStructField"`
	ExportedField        string      `json:"ExportedField"`
	ExportedPointerField *int        `json:"ExportedPointerField"`
	ExportedPointerSlice []int       `json:"ExportedPointerSlice"`
	unexportedField      rune
	unexportedChanField  chan<- string
}

func (t TestStruct1) Method1() int {
	return 0
}

func (t *TestStruct1) Method2(val string) {
	t.ExportedField = "burak"
	return
}

func (t TestStruct1) Method4() {

}

func (t *TestStruct1) Method5() {
}

func (t *TestStruct1) Method6() {
}

type TestStruct2 struct {
	ExportedInnerStructField TestStruct3
}

func (t *TestStruct2) Method7() {
}

type TestStruct3 struct {
	ExportedIntegerField int
}

func TestTypeOfStruct(t *testing.T) {
	typ := TypeOf[TestStruct1]()
	assert.True(t, IsStruct(typ))
	assert.Equal(t, "TestStruct1", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	structType := ToStruct(typ)

	assert.NotNil(t, structType)

	assert.Equal(t, 6, structType.NumMethod())
	assert.Equal(t, 8, structType.NumField())

	testInterface1 := TypeOf[TestInterface1]()
	interfaceType := ToInterface(testInterface1)
	assert.NotNil(t, interfaceType)

	implements := structType.Implements(interfaceType)
	assert.False(t, implements)

	testInterface2 := TypeOf[TestInterface2]()
	interfaceType = ToInterface(testInterface2)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.False(t, implements)

	testInterface3 := TypeOf[TestInterface3]()
	interfaceType = ToInterface(testInterface3)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	testInterface4 := TypeOf[TestInterface4]()
	interfaceType = ToInterface(testInterface4)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.False(t, implements)

	testInterface5 := TypeOf[TestInterface5]()
	interfaceType = ToInterface(testInterface5)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	fields := structType.Fields()
	assert.NotNil(t, fields)

	field := fields[0]
	assert.Equal(t, "TestInterface5", field.Name())
	assert.True(t, field.IsExported())
	assert.True(t, field.IsAnonymous())
	assert.Equal(t, "TestInterface5", field.Type().Name())

	fieldVal, err := field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags := field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists := tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[1]
	assert.Equal(t, "TestEmbeddedStruct", field.Name())
	assert.True(t, field.IsExported())
	assert.True(t, field.IsAnonymous())
	assert.Equal(t, "TestEmbeddedStruct", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[2]
	assert.Equal(t, "ExportedStructField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "TestStruct2", field.Type().Name())

	field = fields[3]
	assert.Equal(t, "ExportedField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "string", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[4]
	assert.Equal(t, "ExportedPointerField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "*int", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[5]
	assert.Equal(t, "ExportedPointerSlice", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "[]int", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[6]
	assert.Equal(t, "unexportedField", field.Name())
	assert.False(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "int32", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 0)
	assert.False(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.False(t, exists)
	assert.Nil(t, tag)

	field = fields[7]
	assert.Equal(t, "unexportedChanField", field.Name())
	assert.False(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "chan<- string", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 0)
	assert.False(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.False(t, exists)
	assert.Nil(t, tag)

	methods := structType.Methods()
	assert.Len(t, methods, 6)

	method := methods[0]
	assert.Equal(t, "Method1", method.Name())
	assert.Equal(t, 0, method.NumParameter())
	assert.Equal(t, 1, method.NumResult())
	assert.True(t, method.IsExported())

	_, err = method.Invoke()
	assert.NotNil(t, err)

	method = methods[1]
	assert.Equal(t, "Method2", method.Name())
	assert.Equal(t, 1, method.NumParameter())
	assert.Equal(t, 0, method.NumResult())
	assert.True(t, method.IsExported())

	assert.Equal(t, "string", method.Parameters()[0].Name())
}

func TestTypeOfStructPointer(t *testing.T) {
	ptrType := TypeOf[*TestStruct1]()
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*TestStruct1", ptr.Name())
	assert.Equal(t, "reflector", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsStruct(typ))
	assert.Equal(t, "TestStruct1", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	structType := ToStruct(typ)

	assert.NotNil(t, structType)

	assert.Equal(t, 6, structType.NumMethod())
	assert.Equal(t, 8, structType.NumField())

	testInterface1 := TypeOf[TestInterface1]()
	interfaceType := ToInterface(testInterface1)
	assert.NotNil(t, interfaceType)

	implements := structType.Implements(interfaceType)
	assert.True(t, implements)

	testInterface2 := TypeOf[TestInterface2]()
	interfaceType = ToInterface(testInterface2)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.False(t, implements)

	testInterface3 := TypeOf[TestInterface3]()
	interfaceType = ToInterface(testInterface3)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	testInterface4 := TypeOf[TestInterface4]()
	interfaceType = ToInterface(testInterface4)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	testInterface5 := TypeOf[TestInterface5]()
	interfaceType = ToInterface(testInterface5)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	fields := structType.Fields()
	assert.NotNil(t, fields)

	field := fields[0]
	assert.Equal(t, "TestInterface5", field.Name())
	assert.True(t, field.IsExported())
	assert.True(t, field.IsAnonymous())
	assert.Equal(t, "TestInterface5", field.Type().Name())

	fieldVal, err := field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags := field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists := tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[1]
	assert.Equal(t, "TestEmbeddedStruct", field.Name())
	assert.True(t, field.IsExported())
	assert.True(t, field.IsAnonymous())
	assert.Equal(t, "TestEmbeddedStruct", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[2]
	assert.Equal(t, "ExportedStructField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "TestStruct2", field.Type().Name())

	field = fields[3]
	assert.Equal(t, "ExportedField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "string", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[4]
	assert.Equal(t, "ExportedPointerField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "*int", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[5]
	assert.Equal(t, "ExportedPointerSlice", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "[]int", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[6]
	assert.Equal(t, "unexportedField", field.Name())
	assert.False(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "int32", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 0)
	assert.False(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.False(t, exists)
	assert.Nil(t, tag)

	field = fields[7]
	assert.Equal(t, "unexportedChanField", field.Name())
	assert.False(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "chan<- string", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 0)
	assert.False(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.False(t, exists)
	assert.Nil(t, tag)

	methods := structType.Methods()
	assert.Len(t, methods, 6)

	method := methods[0]
	assert.Equal(t, "Method1", method.Name())
	assert.Equal(t, 0, method.NumParameter())
	assert.Equal(t, 1, method.NumResult())
	assert.True(t, method.IsExported())

	_, err = method.Invoke()
	assert.NotNil(t, err)

	method = methods[1]
	assert.Equal(t, "Method2", method.Name())
	assert.Equal(t, 1, method.NumParameter())
	assert.Equal(t, 0, method.NumResult())
	assert.True(t, method.IsExported())

	assert.Equal(t, "string", method.Parameters()[0].Name())
}

func TestTypeOfStructObject(t *testing.T) {
	i := -1

	val := TestStruct1{
		TestInterface5: &TestStruct2{},
		TestEmbeddedStruct: TestEmbeddedStruct{
			ExportedAndEmbeddedField:   byte('0'),
			unexportedAndEmbeddedField: 13,
		},
		ExportedField:        "TestValue",
		ExportedPointerField: &i,
		ExportedPointerSlice: []int{1, 3, 5},
		unexportedField:      'A',
		unexportedChanField:  make(chan string),
	}

	typ := TypeOfAny(val)
	assert.True(t, IsStruct(typ))
	assert.Equal(t, "TestStruct1", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	structType := ToStruct(typ)

	assert.NotNil(t, structType)

	assert.Equal(t, 6, structType.NumMethod())
	assert.Equal(t, 8, structType.NumField())

	testInterface1 := TypeOf[TestInterface1]()
	interfaceType := ToInterface(testInterface1)
	assert.NotNil(t, interfaceType)

	implements := structType.Implements(interfaceType)
	assert.False(t, implements)

	testInterface2 := TypeOf[TestInterface2]()
	interfaceType = ToInterface(testInterface2)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.False(t, implements)

	testInterface3 := TypeOf[TestInterface3]()
	interfaceType = ToInterface(testInterface3)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	testInterface4 := TypeOf[TestInterface4]()
	interfaceType = ToInterface(testInterface4)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.False(t, implements)

	testInterface5 := TypeOf[TestInterface5]()
	interfaceType = ToInterface(testInterface5)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	fields := structType.Fields()
	assert.NotNil(t, fields)

	field := fields[0]
	assert.Equal(t, "TestInterface5", field.Name())
	assert.True(t, field.IsExported())
	assert.True(t, field.IsAnonymous())
	assert.Equal(t, "TestInterface5", field.Type().Name())

	fieldVal, err := field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	_, ok := fieldVal.(*TestStruct2)
	assert.True(t, ok)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags := field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists := tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[1]
	assert.Equal(t, "TestEmbeddedStruct", field.Name())
	assert.True(t, field.IsExported())
	assert.True(t, field.IsAnonymous())
	assert.Equal(t, "TestEmbeddedStruct", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	testEmbeddedStructVal, ok := fieldVal.(TestEmbeddedStruct)
	assert.True(t, ok)
	assert.NotNil(t, testEmbeddedStructVal)
	assert.Equal(t, val.TestEmbeddedStruct, testEmbeddedStructVal)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[2]
	assert.Equal(t, "ExportedStructField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "TestStruct2", field.Type().Name())

	field = fields[3]
	assert.Equal(t, "ExportedField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "string", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	assert.Equal(t, val.ExportedField, fieldVal)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[4]
	assert.Equal(t, "ExportedPointerField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "*int", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	assert.Equal(t, val.ExportedPointerField, fieldVal)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[5]
	assert.Equal(t, "ExportedPointerSlice", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "[]int", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	assert.Equal(t, val.ExportedPointerSlice, fieldVal)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[6]
	assert.Equal(t, "unexportedField", field.Name())
	assert.False(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "int32", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 0)
	assert.False(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.False(t, exists)
	assert.Nil(t, tag)

	field = fields[7]
	assert.Equal(t, "unexportedChanField", field.Name())
	assert.False(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "chan<- string", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue("anyTestValue")
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 0)
	assert.False(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.False(t, exists)
	assert.Nil(t, tag)

	methods := structType.Methods()
	assert.Len(t, methods, 6)

	method := methods[0]
	assert.Equal(t, "Method1", method.Name())
	assert.Equal(t, "reflector", method.PackageName())
	assert.Equal(t, "github.com/procyon-projects/reflector", method.PackagePath())
	assert.Equal(t, 0, method.NumParameter())
	assert.Equal(t, 1, method.NumResult())
	assert.True(t, method.IsExported())
	assert.False(t, method.CanSet())

	_, err = method.Invoke()
	assert.Nil(t, err)

	method = methods[1]
	assert.Equal(t, "Method2", method.Name())
	assert.Equal(t, 1, method.NumParameter())
	assert.Equal(t, 0, method.NumResult())
	assert.True(t, method.IsExported())

	assert.Equal(t, "string", method.Parameters()[0].Name())

	_, err = method.Invoke("anyValue")
	assert.Nil(t, err)

}

func TestTypeOfStructObjectPointer(t *testing.T) {
	i := -1

	val := TestStruct1{
		TestInterface5: &TestStruct2{},
		TestEmbeddedStruct: TestEmbeddedStruct{
			ExportedAndEmbeddedField:   byte('0'),
			unexportedAndEmbeddedField: 13,
		},
		ExportedStructField: TestStruct2{
			TestStruct3{
				102,
			},
		},
		ExportedField:        "TestValue",
		ExportedPointerField: &i,
		ExportedPointerSlice: []int{1, 3, 5},
		unexportedField:      'A',
		unexportedChanField:  make(chan string),
	}

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr := ToPointer(ptrType)

	assert.NotNil(t, ptr)

	assert.Equal(t, "*TestStruct1", ptr.Name())
	assert.Equal(t, "reflector", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsStruct(typ))
	assert.Equal(t, "TestStruct1", typ.Name())
	assert.Equal(t, "reflector", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	structType := ToStruct(typ)

	assert.NotNil(t, structType)

	assert.Equal(t, 6, structType.NumMethod())
	assert.Equal(t, 8, structType.NumField())

	testInterface1 := TypeOf[TestInterface1]()
	interfaceType := ToInterface(testInterface1)
	assert.NotNil(t, interfaceType)

	implements := structType.Implements(interfaceType)
	assert.True(t, implements)

	testInterface2 := TypeOf[TestInterface2]()
	interfaceType = ToInterface(testInterface2)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.False(t, implements)

	testInterface3 := TypeOf[TestInterface3]()
	interfaceType = ToInterface(testInterface3)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	testInterface4 := TypeOf[TestInterface4]()
	interfaceType = ToInterface(testInterface4)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	testInterface5 := TypeOf[TestInterface5]()
	interfaceType = ToInterface(testInterface5)
	assert.NotNil(t, interfaceType)

	implements = structType.Implements(interfaceType)
	assert.True(t, implements)

	fields := structType.Fields()
	assert.NotNil(t, fields)

	field := fields[0]
	assert.Equal(t, "TestInterface5", field.Name())
	assert.True(t, field.IsExported())
	assert.True(t, field.IsAnonymous())
	assert.Equal(t, "TestInterface5", field.Type().Name())

	fieldVal, err := field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	_, ok := fieldVal.(*TestStruct2)
	assert.True(t, ok)

	anotherTestStruct := &TestStruct2{}
	err = field.SetValue(anotherTestStruct)
	assert.Nil(t, err)

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)
	assert.Equal(t, anotherTestStruct, fieldVal)

	tags := field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists := tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[1]
	assert.Equal(t, "TestEmbeddedStruct", field.Name())
	assert.True(t, field.IsExported())
	assert.True(t, field.IsAnonymous())
	assert.Equal(t, "TestEmbeddedStruct", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	testEmbeddedStructVal, ok := fieldVal.(TestEmbeddedStruct)
	assert.True(t, ok)
	assert.NotNil(t, testEmbeddedStructVal)
	assert.Equal(t, val.TestEmbeddedStruct, testEmbeddedStructVal)

	anotherEmbeddedStruct := TestEmbeddedStruct{
		ExportedAndEmbeddedField:   6,
		unexportedAndEmbeddedField: 1,
	}
	err = field.SetValue(anotherEmbeddedStruct)
	assert.Nil(t, err)

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)
	assert.Equal(t, anotherEmbeddedStruct, fieldVal)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[2]

	assert.Equal(t, "ExportedStructField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "TestStruct2", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	testStruct2Val, ok := fieldVal.(TestStruct2)
	assert.True(t, ok)
	assert.NotNil(t, testStruct2Val)
	assert.Equal(t, val.ExportedStructField, testStruct2Val)

	fieldType := field.Type()
	testStruct2Type := ToStruct(fieldType)
	assert.NotNil(t, testStruct2Type)

	fieldVal, err = testStruct2Type.Value()
	assert.Nil(t, err)
	assert.Equal(t, val.ExportedStructField, testStruct2Val)

	innerField := testStruct2Type.Fields()[0]
	innerFieldVal, err := innerField.Value()
	assert.Nil(t, err)
	assert.Equal(t, val.ExportedStructField.ExportedInnerStructField, innerFieldVal)

	anotherTestStruct3 := TestStruct3{ExportedIntegerField: 111}
	err = innerField.SetValue(anotherTestStruct3)
	assert.Nil(t, err)

	innerFieldVal, err = innerField.Value()
	assert.Nil(t, err)
	assert.Equal(t, anotherTestStruct3, innerFieldVal)

	anotherTestStruct2Val := TestStruct2{TestStruct3{ExportedIntegerField: 6}}
	err = field.SetValue(anotherTestStruct2Val)
	assert.Nil(t, err)

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)
	assert.Equal(t, anotherTestStruct2Val, fieldVal)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[3]
	assert.Equal(t, "ExportedField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "string", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	assert.Equal(t, val.ExportedField, fieldVal)

	err = field.SetValue("anyTestValue")
	assert.Nil(t, err)

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)
	assert.Equal(t, "anyTestValue", fieldVal)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[4]
	assert.Equal(t, "ExportedPointerField", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "*int", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	assert.Equal(t, val.ExportedPointerField, fieldVal)

	anotherInt := 29
	err = field.SetValue(&anotherInt)
	assert.Nil(t, err)

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)
	assert.Equal(t, &anotherInt, fieldVal)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[5]
	assert.Equal(t, "ExportedPointerSlice", field.Name())
	assert.True(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "[]int", field.Type().Name())

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)

	assert.Equal(t, val.ExportedPointerSlice, fieldVal)

	anotherSlice := []int{7, 1, 8}
	err = field.SetValue(anotherSlice)
	assert.Nil(t, err)

	fieldVal, err = field.Value()
	assert.NotNil(t, fieldVal)
	assert.Nil(t, err)
	assert.Equal(t, anotherSlice, fieldVal)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.True(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.True(t, exists)
	assert.NotNil(t, tag)
	assert.Equal(t, "json", tag.Name())
	assert.Equal(t, field.Name(), tag.Value())

	field = fields[6]
	assert.Equal(t, "unexportedField", field.Name())
	assert.False(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "int32", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue(32)
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 0)
	assert.False(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.False(t, exists)
	assert.Nil(t, tag)

	field = fields[7]
	assert.Equal(t, "unexportedChanField", field.Name())
	assert.False(t, field.IsExported())
	assert.False(t, field.IsAnonymous())
	assert.Equal(t, "chan<- string", field.Type().Name())

	fieldVal, err = field.Value()
	assert.Nil(t, fieldVal)
	assert.NotNil(t, err)

	err = field.SetValue(make(chan<- string))
	assert.NotNil(t, err)

	tags = field.Tags()
	assert.NotNil(t, tags)
	assert.Len(t, tags, 0)
	assert.False(t, tags.Contains("json"))
	tag, exists = tags.Find("json")
	assert.False(t, exists)
	assert.Nil(t, tag)

	field, exists = structType.FieldByName("ExportedField")
	assert.True(t, exists)
	assert.NotNil(t, field)

	err = field.SetValue("FieldByName")
	assert.Nil(t, err)

	fieldVal, err = field.Value()
	assert.Equal(t, "FieldByName", fieldVal)
	assert.Nil(t, err)

	field, exists = structType.FieldByName("NotExist")
	assert.False(t, exists)
	assert.Nil(t, field)

	field, exists = structType.Field(3)
	assert.True(t, exists)
	assert.NotNil(t, field)

	err = field.SetValue("FieldByIndex")
	assert.Nil(t, err)

	fieldVal, err = field.Value()
	assert.Equal(t, "FieldByIndex", fieldVal)
	assert.Nil(t, err)

	methods := structType.Methods()
	assert.Len(t, methods, 6)

	method := methods[0]
	assert.Equal(t, "Method1", method.Name())
	assert.Equal(t, "reflector", method.PackageName())
	assert.Equal(t, "github.com/procyon-projects/reflector", method.PackagePath())
	assert.Equal(t, 0, method.NumParameter())
	assert.Equal(t, 1, method.NumResult())
	assert.True(t, method.IsExported())
	assert.False(t, method.CanSet())

	method = methods[1]
	assert.Equal(t, "Method2", method.Name())
	assert.Equal(t, 1, method.NumParameter())
	assert.Equal(t, 0, method.NumResult())
	assert.True(t, method.IsExported())

	assert.Equal(t, "string", method.Parameters()[0].Name())

	_, err = method.Invoke("anyValue")
	assert.Nil(t, err)
}
