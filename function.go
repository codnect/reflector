package reflector

import (
	"reflect"
	"runtime"
	"strings"
)

type Function interface {
	Receiver() (Type, bool)
	HasReceiver() bool
	Parameters() []Type
	NumParameter() int
	Results() []Type
	NumResult() int
}

type functionType struct {
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (f *functionType) Name() string {
	name := runtime.FuncForPC(f.reflectValue.Pointer()).Name()
	dotLastIndex := strings.LastIndex(name, ".")

	if dotLastIndex != -1 {
		return name[dotLastIndex+1:]
	}

	return name
}

func (f *functionType) PackageName() string {
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

func (f *functionType) ReflectType() reflect.Type {
	return f.reflectType
}

func (f *functionType) ReflectValue() *reflect.Value {
	return f.reflectValue
}

func (f *functionType) Receiver() (Type, bool) {
	return nil, false
}

func (f *functionType) HasReceiver() bool {
	return false
}

func (f *functionType) Parameters() []Type {
	return nil
}

func (f *functionType) NumParameter() int {
	return 0
}

func (f *functionType) Results() []Type {
	return nil
}

func (f *functionType) NumResult() int {
	return 0
}
