package main

import (
	"fmt"
	"os"
	"strconv"
)

func calc(arg int) int {
	if arg < 1 {
		panic("I need a positive number where (n > 1)")
	}

	if arg == 1 {
		return 0
	}
	if arg == 2 {
		return 1
	}

	a := 0
	b := 1
	var c = a + b

	for i := 0; i < arg-2; i += 1 {
		c = a + b
		a = b
		b = c
	}
	return c
}
func main() {

	if len(os.Args) < 2 {
		panic("I need a number, write it as argument")
	}

	arg, _ := strconv.ParseInt(os.Args[1], 10, 32)

	fmt.Println(arg, "number of fibonachi is: ", calc(int(arg)))
}
