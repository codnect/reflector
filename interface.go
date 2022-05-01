package reflector

import (
	"reflect"
	"strings"
)

type Interface interface {
	Type
	Elem() Type
	Methods() []Function
	NumMethod() int
}

type interfaceType struct {
	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (i *interfaceType) Name() string {
	if i.reflectType.Name() == "" {
		return "any"
	}

	return i.reflectType.Name()
}

func (i *interfaceType) PackageName() string {
	name := i.reflectType.PkgPath()
	slashLastIndex := strings.LastIndex(name, "/")

	if slashLastIndex != -1 {
		name = name[slashLastIndex+1:]
	}

	return name
}

func (i *interfaceType) PackagePath() string {
	return i.reflectType.PkgPath()
}

func (i *interfaceType) HasValue() bool {
	return i.reflectValue != nil
}

func (i *interfaceType) Parent() Type {
	return i.parent
}

func (i *interfaceType) ReflectType() reflect.Type {
	return i.reflectType
}

func (i *interfaceType) ReflectValue() *reflect.Value {
	return i.reflectValue
}

func (i *interfaceType) Elem() Type {
	return nil
}

func (i *interfaceType) Methods() []Function {
	functions := make([]Function, 0)
	numMethod := i.reflectType.NumMethod()

	for index := 0; index < numMethod; index++ {
		function := i.reflectType.Method(index)
		functions = append(functions, &functionType{
			name:        function.Name,
			pkgPath:     function.PkgPath,
			isExported:  function.IsExported(),
			reflectType: function.Type,
		})
	}

	return functions
}

func (i *interfaceType) NumMethod() int {
	return i.reflectType.NumMethod()
}
