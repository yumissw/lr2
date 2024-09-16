package main

import "fmt"

func func1(x int) (s string) {

	if x == 0 {
		s = "zero"
	} else if x > 0 {
		s = "positive"
	} else if x < 0 {
		s = "negative"
	}

	return s
}

func main() {

	var x int

	fmt.Print("введите число: ")
	fmt.Scan(&x)

	fmt.Println(func1(x))
}
