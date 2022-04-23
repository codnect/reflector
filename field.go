package reflector

type Field interface {
	Type() Type
	CanSet() bool
	Value() any
	SetValue(value any)
	Tags() []Tag
}

type field struct {
}

func (f *field) Type() Type {
	return nil
}

func (f *field) CanSet() bool {
	return false
}

func (f *field) Value() any {
	return nil
}

func (f *field) Tags() []Tag {
	return nil
}

func (f *field) SetValue(value any) {
}
