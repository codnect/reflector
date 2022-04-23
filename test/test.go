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
	x := reflector.TypeOf[Test[string]]()

	if str, ok := reflector.ToInterface(x); ok {

		if str.HasReference() {

		}

		fields := str.Methods()

		if fields != nil {

		}
	}

	if name == nil {

	}
}
