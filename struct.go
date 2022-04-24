package reflector

import (
	"reflect"
	"strings"
)

type Struct interface {
	Type
	Instantiable
	CanSet() bool
	Fields() []Field
	NumField() int
	Methods() []Function
	NumMethod() int
	Implements(i Interface) bool
}

type structType struct {
	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *structType) Name() string {
	return s.reflectType.Name()
}

func (s *structType) PackageName() string {
	name := s.reflectType.PkgPath()
	slashLastIndex := strings.LastIndex(name, "/")

	if slashLastIndex != -1 {
		name = name[slashLastIndex+1:]
	}

	return name
}

func (s *structType) HasValue() bool {
	return s.reflectValue != nil
}

func (s *structType) Parent() Type {
	return s.parent
}

func (s *structType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *structType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *structType) NumField() int {
	return s.reflectType.NumField()
}

func (s *structType) CanSet() bool {
	if s.reflectValue == nil {
		return false
	}

	return s.reflectValue.CanSet()
}

func (s *structType) Fields() []Field {
	fields := make([]Field, 0)

	numField := s.reflectType.NumField()
	for i := 0; i < numField; i++ {
		f := s.reflectType.Field(i)
		fields = append(fields, &field{
			index:       i,
			structType:  s,
			structField: f,
		})
	}
	return fields
}

func (s *structType) Methods() []Function {
	functions := make([]Function, 0)

	reflectType := s.reflectType

	if s.Parent() != nil {
		reflectType = s.Parent().ReflectType()
	}

	numMethod := reflectType.NumMethod()
	for i := 0; i < numMethod; i++ {
		function := reflectType.Method(i)
		functions = append(functions, &functionType{
			name:        function.Name,
			pkgPath:     function.PkgPath,
			isExported:  function.IsExported(),
			reflectType: function.Type,
		})
	}

	return functions
}

func (s *structType) NumMethod() int {
	if s.Parent() != nil {
		return s.Parent().ReflectType().NumMethod()
	}

	return s.reflectType.NumMethod()
}

func (s *structType) Implements(i Interface) bool {
	if s.Parent() != nil {
		return s.Parent().ReflectType().Implements(i.ReflectType())
	}

	return s.reflectType.Implements(i.ReflectType())
}

func (s *structType) Instantiate() Value {
	return &value{
		reflect.New(s.reflectType),
	}
}
