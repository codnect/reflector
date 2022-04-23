package reflector

import "reflect"

type Interface interface {
	Type
	Methods() []Function
	NumMethod() int
}

type interfaceType struct {
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (i *interfaceType) Name() string {
	return i.reflectType.Name()
}

func (i *interfaceType) PackageName() string {
	return i.reflectType.Name()
}

func (i *interfaceType) HasReference() bool {
	return i.reflectValue != nil
}

func (i *interfaceType) ReflectType() reflect.Type {
	return i.reflectType
}

func (i *interfaceType) ReflectValue() *reflect.Value {
	return i.reflectValue
}

func (i *interfaceType) Methods() []Function {
	functions := make([]Function, 0)
	numMethod := i.reflectType.NumMethod()

	for index := 0; index < numMethod; index++ {
		function := i.reflectType.Method(index)
		if function.IsExported() {

		}
		functions = append(functions, &functionType{})
	}

	return functions
}

func (i *interfaceType) NumMethod() int {
	return i.reflectType.NumMethod()
}
