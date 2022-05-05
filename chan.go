package reflector

import (
	"errors"
	"reflect"
	"strings"
)

type ChanDirection int

const (
	RECEIVE ChanDirection    = 1 << iota // <-chan
	SEND                                 // chan<-
	BOTH    = RECEIVE | SEND             // chan
)

type Chan interface {
	Type
	Direction() ChanDirection
	Elem() Type
	Cap() (int, error)
	Send(value any) error
	Receive() (any, error)
	TrySend(value any) error
	TryReceive() (any, error)
}

type chanType struct {
	parent       Type
	elem         Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (c *chanType) Name() string {
	var builder strings.Builder

	if c.reflectType.ChanDir() == reflect.RecvDir {
		builder.WriteString("<-")
	}

	builder.WriteString("chan")

	if c.reflectType.ChanDir() == reflect.SendDir {
		builder.WriteString("<-")
	}

	builder.WriteString(" ")
	builder.WriteString(c.elem.Name())

	return builder.String()
}

func (c *chanType) PackageName() string {
	return ""
}

func (c *chanType) PackagePath() string {
	return ""
}

func (c *chanType) CanSet() bool {
	if c.reflectValue == nil {
		return false
	}

	return c.reflectValue.CanSet()
}

func (c *chanType) HasValue() bool {
	return c.reflectValue != nil
}

func (c *chanType) Value() (any, error) {
	if c.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return c.reflectValue.Interface(), nil
}

func (c *chanType) SetValue(val any) error {
	if !c.CanSet() {
		return errors.New("value cannot be set")
	}

	c.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (c *chanType) Parent() Type {
	return c.parent
}

func (c *chanType) ReflectType() reflect.Type {
	return c.reflectType
}

func (c *chanType) ReflectValue() *reflect.Value {
	return c.reflectValue
}

func (c *chanType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return c.reflectType == another.ReflectType()
}

func (c *chanType) IsInstantiable() bool {
	return true
}

func (c *chanType) Instantiate() (Value, error) {
	return &value{
		reflect.New(c.reflectType),
	}, nil
}

func (c *chanType) Send(value any) error {
	if c.reflectValue == nil {
		return errors.New("value reference is nil")
	}

	c.reflectValue.Send(reflect.ValueOf(value))
	return nil
}

func (c *chanType) Receive() (any, error) {
	if c.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	val, ok := c.reflectValue.Recv()

	if !ok {
		return nil, errors.New("value could not be received")
	}

	return val.Interface(), nil
}

func (c *chanType) TrySend(value any) error {
	if c.reflectValue == nil {
		return errors.New("value reference is nil")
	}

	if ok := c.reflectValue.TrySend(reflect.ValueOf(value)); !ok {
		return errors.New("value could not be sent")
	}

	return nil
}

func (c *chanType) TryReceive() (any, error) {
	if c.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	val, ok := c.reflectValue.TryRecv()

	if !ok {
		return nil, errors.New("value could not be received")
	}

	return val.Interface(), nil
}

func (c *chanType) Direction() ChanDirection {
	switch c.reflectType.ChanDir() {
	case reflect.RecvDir:
		return RECEIVE
	case reflect.SendDir:
		return SEND
	default:
		return BOTH
	}
}

func (c *chanType) Elem() Type {
	return c.elem
}

func (c *chanType) Cap() (int, error) {
	if c.reflectValue == nil {
		return -1, errors.New("value reference is nil")
	}

	return c.reflectValue.Cap(), nil
}
