package reflector

import (
	"errors"
	"reflect"
	"strconv"
)

type Field interface {
	Name() string
	IsExported() bool
	IsAnonymous() bool
	Type() Type
	CanSet() bool
	Value() (any, error)
	SetValue(value any) error
	Tags() Tags
}

type field struct {
	index       int
	structType  *structType
	structField reflect.StructField
}

func (f *field) Name() string {
	return f.structField.Name
}

func (f *field) IsExported() bool {
	return f.structField.IsExported()
}

func (f *field) IsAnonymous() bool {
	return f.structField.Anonymous
}

func (f *field) Type() Type {
	v := f.structType.reflectValue.Field(f.index)
	return typeOf(f.structField.Type, &v, f.structType)
}

func (f *field) CanSet() bool {
	return f.structType.CanSet()
}

func (f *field) Value() (any, error) {
	if f.structType.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	return f.structType.reflectValue.Field(f.index).Interface(), nil
}

func (f *field) SetValue(value any) error {
	if !f.CanSet() {
		return errors.New("value cannot be set")
	}

	f.structType.reflectValue.Field(f.index).Set(reflect.ValueOf(value))
	return nil
}

func (f *field) Tags() Tags {
	fieldTags := make([]Tag, 0)

	tags := f.structField.Tag
	for tags != "" {
		i := 0
		for i < len(tags) && tags[i] == ' ' {
			i++
		}
		tags = tags[i:]
		if tags == "" {
			break
		}

		i = 0
		for i < len(tags) && tags[i] > ' ' && tags[i] != ':' && tags[i] != '"' && tags[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tags) || tags[i] != ':' || tags[i+1] != '"' {
			break
		}
		name := string(tags[:i])
		tags = tags[i+1:]

		i = 1
		for i < len(tags) && tags[i] != '"' {
			if tags[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tags) {
			break
		}
		quotedValue := string(tags[:i+1])
		tags = tags[i+1:]

		value, err := strconv.Unquote(quotedValue)
		if err != nil {
			break
		}

		fieldTag := &tag{name, value}
		fieldTags = append(fieldTags, fieldTag)
	}

	return fieldTags
}
