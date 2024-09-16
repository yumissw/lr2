package main

import (
	"fmt"
)

func func1(a int, b int) (x int) {
	x = (a + b) / 2

	return x
}

func main() {
	var a, b, x int

	fmt.Print("введите a: ")
	fmt.Scan(&a)

	fmt.Print("введите b: ")
	fmt.Scan(&b)

	x = func1(a, b)

	fmt.Println("среднее значение = ", x)
}
