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
	name := [5]string{"hello", "world!"}
	x := reflector.TypeOfAny(&name)

	arr, _ := reflector.ToArray(x)
	//subArr, _ := reflector.ToArray(arr.Elem())
	//m := subArr.Instantiate().Elem().([2]int)

	arr.Set(0, "burak")

	/*if len(m) == 9 {

	}*/

	if arr == nil {

	}

	if str, ok := reflector.ToStruct(x); ok {

		if str.HasValue() {

		}

		fields := str.Methods()

		if fields != nil {

		}
	}

}
