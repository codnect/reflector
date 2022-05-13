package reflector

import (
	"errors"
	"reflect"
	"strings"
)

type Custom interface {
	Type
	Underlying() Type
	Methods() []Function
	NumMethod() int
	Implements(i Interface) bool
}

type customType struct {
	parent         Type
	reflectType    reflect.Type
	reflectValue   *reflect.Value
	underlyingType Type
}

func (c *customType) Name() string {
	return c.reflectType.Name()
}

func (c *customType) PackageName() string {
	name := c.reflectType.PkgPath()
	slashLastIndex := strings.LastIndex(name, "/")

	if slashLastIndex != -1 {
		name = name[slashLastIndex+1:]
	}

	return name
}

func (c *customType) PackagePath() string {
	return c.reflectType.PkgPath()
}

func (c *customType) CanSet() bool {
	if c.reflectValue == nil {
		return false
	}

	return c.reflectValue.CanSet()
}

func (c *customType) HasValue() bool {
	return c.reflectValue != nil
}

func (c *customType) Value() (any, error) {
	if c.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return c.reflectValue.Interface(), nil
}

func (c *customType) SetValue(val any) error {
	if !c.CanSet() {
		return errors.New("value cannot be set")
	}

	c.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (c *customType) Parent() Type {
	return c.parent
}

func (c *customType) ReflectType() reflect.Type {
	return c.reflectType
}

func (c *customType) ReflectValue() *reflect.Value {
	return c.reflectValue
}

func (c *customType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return c.reflectType == another.ReflectType()
}

func (c *customType) IsInstantiable() bool {
	return false
}

func (c *customType) Instantiate() (Value, error) {
	return &value{
		reflect.New(c.reflectType),
	}, nil
}

func (c *customType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if c.reflectValue == nil {
		return c.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return c.reflectValue.CanConvert(typ.ReflectType())
}

func (c *customType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if c.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !c.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := c.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (c *customType) Underlying() Type {
	return nil
}

func (c *customType) Methods() []Function {
	functions := make([]Function, 0)

	reflectType := c.reflectType

	if c.Parent() != nil {
		reflectType = c.Parent().ReflectType()
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

func (c *customType) NumMethod() int {
	if c.Parent() != nil {
		return c.Parent().ReflectType().NumMethod()
	}

	return c.reflectType.NumMethod()
}

func (c *customType) Implements(i Interface) bool {
	if c.Parent() != nil {
		return c.Parent().ReflectType().Implements(i.ReflectType())
	}

	return c.reflectType.Implements(i.ReflectType())
}
