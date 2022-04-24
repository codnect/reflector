package reflector

import (
	"reflect"
	"runtime"
	"strings"
)

type Function interface {
	Type
	IsExported() bool
	Receiver() (Type, bool)
	HasReceiver() bool
	Parameters() []Type
	NumParameter() int
	Results() []Type
	NumResult() int
}

type functionType struct {
	name       string
	pkgPath    string
	isExported bool
	receiver   Type

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (f *functionType) Name() string {
	if f.name != "" {
		return f.name
	}

	name := runtime.FuncForPC(f.reflectValue.Pointer()).Name()
	dotLastIndex := strings.LastIndex(name, ".")

	if dotLastIndex != -1 {
		return name[dotLastIndex+1:]
	}

	return name
}

func (f *functionType) PackageName() string {
	if f.pkgPath != "" {
		name := f.pkgPath
		slashLastIndex := strings.LastIndex(name, "/")

		if slashLastIndex != -1 {
			name = name[slashLastIndex+1:]
		}

		return name
	}

	name := runtime.FuncForPC(f.reflectValue.Pointer()).Name()
	dotLastIndex := strings.LastIndex(name, ".")

	if dotLastIndex != -1 {
		name = name[:dotLastIndex]
	}

	slashLastIndex := strings.LastIndex(name, "/")

	if slashLastIndex != -1 {
		name = name[slashLastIndex+1:]
	}

	return name
}

func (f *functionType) HasValue() bool {
	return f.reflectValue != nil
}

func (f *functionType) Parent() Type {
	return f.parent
}

func (f *functionType) ReflectType() reflect.Type {
	return f.reflectType
}

func (f *functionType) ReflectValue() *reflect.Value {
	return f.reflectValue
}

func (f *functionType) IsExported() bool {
	return f.isExported
}

func (f *functionType) Receiver() (Type, bool) {
	return nil, false
}

func (f *functionType) HasReceiver() bool {
	return f.receiver != nil
}

func (f *functionType) Parameters() []Type {
	parameters := make([]Type, 0)
	numIn := f.reflectType.NumIn()

	for index := 0; index < numIn; index++ {
		typ := f.reflectType.In(index)
		parameters = append(parameters, typeOf(typ, nil, nil))
	}
	return parameters
}

func (f *functionType) NumParameter() int {
	return f.reflectType.NumIn()
}

func (f *functionType) Results() []Type {
	results := make([]Type, 0)
	numOut := f.reflectType.NumOut()

	for index := 0; index < numOut; index++ {
		typ := f.reflectType.Out(index)
		results = append(results, typeOf(typ, nil, nil))
	}
	return results
}

func (f *functionType) NumResult() int {
	return f.reflectType.NumOut()
}
