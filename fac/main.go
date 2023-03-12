package main

import (
	"fmt"
	"os"
	"strconv"
)

func calc(num int) int {
	a := 1
	for i := 1; i <= num; i += 1 {
		a *= i
	}
	return a
}

func main() {
	if len(os.Args) < 2 {
		panic("I need number to calculate factorial from zero to given number")
	}

	arg, _ := strconv.ParseInt(os.Args[1], 10, 32)

	fmt.Println("answer is: ", calc(int(arg)))

}
