package reflector

type Array interface {
	Elem() Type
}

type arrayType struct {
	elem Type
}

func (a *arrayType) Elem() Type {
	return nil
}

func (a *arrayType) Instantiate() any {
	return nil
}
