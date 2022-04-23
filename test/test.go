package main

import (
	"github.com/procyon-projects/reflector"
)

type Person struct {
	name string
}

type Test[T any] interface {
	Find()
}

func main() {
	name := map[string]string{}
	x := reflector.TypeOfAny(Person{})

	h := x.Name()
	h = x.PackageName()

	if h == "" {

	}

	if str, ok := reflector.ToStruct(x); ok {

		if str.HasReference() {

		}

		fields := str.Methods()

		if fields != nil {

		}
	}

	if name == nil {

	}
}
