package reflector

import (
	"errors"
	"reflect"
	"strings"
)

type Struct interface {
	Type
	Fields() []Field
	Field(index int) (Field, bool)
	FieldByName(name string) (Field, bool)
	NumField() int
	Methods() []Method
	Method(index int) (Method, bool)
	MethodByName(name string) (Method, bool)
	NumMethod() int
	Implements(i Interface) bool
	Embeds(another Type) bool
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
	if another == nil {
		return false
	}

	return s.reflectType == another.ReflectType()
}

func (s *structType) IsInstantiable() bool {
	return true
}

func (s *structType) Instantiate() (Value, error) {
	return &value{
		reflect.New(s.reflectType),
	}, nil
}

func (s *structType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if s.reflectValue == nil {
		return s.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return s.reflectValue.CanConvert(typ.ReflectType())
}

func (s *structType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if s.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !s.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := s.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (s *structType) NumField() int {
	return s.reflectType.NumField()
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
	if index < 0 || index >= s.NumField() {
		return nil, false
	}

	structField := s.reflectType.Field(index)

	return &field{
		index:       index,
		structType:  s,
		structField: structField,
	}, true
}

func (s *structType) FieldByName(name string) (Field, bool) {
	structField, exists := s.reflectType.FieldByName(name)

	if !exists {
		return nil, false
	}

	return &field{
		index:       structField.Index[0],
		structType:  s,
		structField: structField,
	}, true
}

func (s *structType) Methods() []Method {
	methods := make([]Method, 0)

	typ := s.nilType
	if s.Parent() != nil {
		typ = s.nilType.Elem()
	}

	numMethod := s.NumMethod()

	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)

		methods = append(methods, &methodType{
			parent:        s,
			reflectMethod: method,
		})
	}

	return methods
}

func (s *structType) Method(index int) (Method, bool) {
	if index < 0 || index >= s.NumMethod() {
		return nil, false
	}

	typ := s.nilType
	if s.Parent() != nil {
		typ = s.nilType.Elem()
	}

	method := typ.Method(index)

	return &methodType{
		parent:        s,
		reflectMethod: method,
	}, true
}

func (s *structType) MethodByName(name string) (Method, bool) {
	typ := s.nilType
	if s.Parent() != nil {
		typ = s.nilType.Elem()
	}

	method, exists := typ.MethodByName(name)

	if !exists {
		return nil, false
	}

	return &methodType{
		parent:        s,
		reflectMethod: method,
	}, true
}

func (s *structType) NumMethod() int {
	if s.Parent() != nil {
		return s.nilType.Elem().NumMethod()
	}

	return s.nilType.NumMethod()
}

func (s *structType) Implements(i Interface) bool {
	if i == nil {
		return false
	}

	if s.Parent() != nil {
		return s.Parent().ReflectType().Implements(i.ReflectType())
	}

	return s.reflectType.Implements(i.ReflectType())
}

func (s *structType) Embeds(another Type) bool {
	if another == nil {
		return false
	}

	visitedMap := make(map[string]bool, 0)
	return s.embeds(another, visitedMap)
}

func (s *structType) embeds(candidate Type, visitedMap map[string]bool) bool {

	for _, field := range s.Fields() {
		if field.IsAnonymous() {
			fieldType := field.Type()

			if visitedMap[fieldType.PackagePath()+"@"+fieldType.PackageName()] {
				continue
			}

			visitedMap[fieldType.PackagePath()+"@"+fieldType.PackageName()] = true

			if candidate.Compare(fieldType) {
				return true
			}

			if IsStruct(fieldType) {
				structType := ToStruct(fieldType)

				if structType.NumField() == 0 {
					continue
				}

				returnValue := structType.Embeds(candidate)

				if returnValue {
					return true
				}
			}

			if IsInterface(fieldType) {
				// TODO complete method
				interfaceType := ToInterface(fieldType)
				interfaceType.Methods()
			}
		}
	}

	return false
}
