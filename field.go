package reflector

type Field interface {
	Type() Type
	CanSet() bool
	Value() any
	SetValue(value any)
	Tags() Tags
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

func (f *field) Tags() Tags {
	return nil
}

func (f *field) SetValue(value any) {
}
