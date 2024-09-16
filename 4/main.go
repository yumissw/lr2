package main

import (
	"fmt"
)

func func1(s string) (x int) {
	x = len(s)

	return x
}

func main() {

	var s string
	fmt.Print("введите строку: ")

	var x int
	x = func1(s)

	fmt.Println("длина строки: ", x)

}
