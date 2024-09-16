package main

import (
	"fmt"
)

type Rectangle struct {
	a int
	b int
}

func func1(a int, b int) (s int) {
	s = a * b

	return s
}

func main() {

	var x Rectangle = Rectangle{a: 10, b: 3}

	var s int

	s = func1(x.a, x.b)

	fmt.Println("площадь: ", s)

}
