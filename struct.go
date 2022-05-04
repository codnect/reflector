package reflector

import (
	"errors"
	"reflect"
	"strings"
)

type Struct interface {
	Type
	Instantiable
	CanSet() bool
	Value() (any, error)
	SetValue(val any) error
	Fields() []Field
	Field(index int) (Field, bool)
	FieldByName(name string) (Field, bool)
	NumField() int
	Methods() []Function
	Method(index int) (Function, bool)
	MethodByName(name string) (Function, bool)
	NumMethod() int
	Implements(i Interface) (bool, error)
	Embeds(another Type) (bool, error)
}

type structType struct {
	parent       Type
	nilType      reflect.Type
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

func (s *structType) PackagePath() string {
	return s.reflectType.PkgPath()
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

func (s *structType) Compare(another Type) bool {
	return false
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

func (s *structType) Value() (any, error) {
	if s.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return s.reflectValue.Interface(), nil
}

func (s *structType) SetValue(val any) error {
	if !s.CanSet() {
		return errors.New("value cannot be set")
	}

	s.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (s *structType) Fields() []Field {
	fields := make([]Field, 0)

	numField := s.reflectType.NumField()

	for i := 0; i < numField; i++ {
		structField := s.reflectType.Field(i)
		fields = append(fields, &field{
			index:       i,
			structType:  s,
			structField: structField,
		})
	}

	return fields
}

func (s *structType) Field(index int) (Field, bool) {
	return nil, false
}

func (s *structType) FieldByName(name string) (Field, bool) {
	return nil, false
}

func (s *structType) Methods() []Function {
	functions := make([]Function, 0)

	numMethod := s.nilType.NumMethod()
	for i := 0; i < numMethod; i++ {
		function := s.nilType.Method(i)
		functions = append(functions, &functionType{
			name:        function.Name,
			pkgPath:     function.PkgPath,
			isExported:  function.IsExported(),
			reflectType: function.Type,
		})
	}

	return functions
}

func (s *structType) Method(index int) (Function, bool) {
	return nil, false
}

func (s *structType) MethodByName(name string) (Function, bool) {
	return nil, false
}

func (s *structType) NumMethod() int {
	if s.Parent() != nil {
		return s.nilType.Elem().NumMethod()
	}

	return s.nilType.NumMethod()
}

func (s *structType) Implements(i Interface) (bool, error) {
	if i == nil {
		return false, errors.New("given interface cannot be nil")
	}

	if s.Parent() != nil {
		return s.Parent().ReflectType().Implements(i.ReflectType()), nil
	}

	return s.reflectType.Implements(i.ReflectType()), nil
}

func (s *structType) Embeds(another Type) (bool, error) {
	if another == nil {
		return false, errors.New("another cannot be nil")
	}

	visitedMap := make(map[string]bool, 0)
	return s.embeds(another, visitedMap)
}

func (s *structType) embeds(candidate Type, visitedMap map[string]bool) (bool, error) {

	for _, field := range s.Fields() {
		if field.IsAnonymous() {
			fieldType := field.Type()

			if visitedMap[fieldType.PackagePath()+"@"+fieldType.PackageName()] {
				continue
			}

			visitedMap[fieldType.PackagePath()+"@"+fieldType.PackageName()] = true

			if candidate.Compare(fieldType) {
				return true, nil
			}

			structType, isStruct := ToStruct(fieldType)

			if isStruct {
				if structType.NumField() == 0 {
					continue
				}

				returnValue, err := structType.Embeds(candidate)

				if err != nil {
					return false, err
				}

				if returnValue {
					return true, nil
				}
			}

			interfaceType, isInterface := ToInterface(fieldType)

			if isInterface {
				interfaceType.Methods()
			}
		}
	}

	return false, nil
}

func (s *structType) Instantiate() Value {
	return &value{
		reflect.New(s.reflectType),
	}
}
