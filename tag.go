package reflector

type Tags []Tag

func (t Tags) Contains(name string) bool {
	for _, i := range t {
		if i.Name() == name {
			return true
		}
	}

	return false
}

func (t Tags) Find(name string) (Tag, bool) {
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

type tag struct {
	name  string
	value string
}

func (t *tag) Name() string {
	return t.name
}

func (t *tag) Value() string {
	return t.value
}
