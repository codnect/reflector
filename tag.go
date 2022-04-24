package reflector

type Tags []Tag

func (t Tags) FindTag(name string) (Tag, bool) {
	for _, i := range t {
		if i.Name() == name {
			return i, true
		}
	}

	return nil, false
}

type Tag interface {
	Name() string
	Value() string
}
