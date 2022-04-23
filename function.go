package reflector

type Function interface {
	Parameters() []Type
	ParameterCount() int
	Results() []Type
	ResultCount() int
}

type functionType struct {
}

func (f *functionType) Parameters() []Type {
	return nil
}

func (f *functionType) ParameterCount() int {
	return 0
}

func (f *functionType) Results() []Type {
	return nil
}

func (f *functionType) ResultCount() int {
	return 0
}
