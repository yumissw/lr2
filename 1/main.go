package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Print("введите число: ")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("число четное")
	} else {
		fmt.Println("число нечетное")
	}
}
