package main

import (
	"bufio"
	"fmt"
	"os"
)

func func1(s string) int {
	return len(s) - 2
}

func main() {

	var s string
	r := bufio.NewReader(os.Stdin)

	fmt.Print("введите строку: ")
	s, _ = r.ReadString('\n')

	x := func1(s)

	fmt.Println("длина строки: ", x)

}
