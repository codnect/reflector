package main

import (
	"github.com/procyon-projects/reflector"
)

type Person struct {
	name string
}

func main() {
	name := map[string]string{}
	x := reflector.TypeOfAny(name)

	if str, ok := reflector.ToMap(x); ok {

		str.Put("burak", "hello")

		str.Contains("burak")

		keys := str.KeySet()
		values := str.EntrySet()

		if keys == nil {

		}

		if values == nil {

		}

		len := str.Len()

		if len == 1 {

		}
	}

	if name == nil {

	}
}
