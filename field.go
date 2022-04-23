package reflector

type Field interface {
	Type() Type
	CanSet() bool
	Value() any
	SetValue(value any)
	Tags() []Tag
}
