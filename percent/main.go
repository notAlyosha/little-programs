package main

import (
	"fmt"
	"os"
	"strconv"
)

func calc(percent, ofPercent float64) float64 {
	return 100 * ((percent / 100) * (ofPercent / 100))
}

func main() {
	if len(os.Args) < 3 {
		panic("I need 2 argument where first is percent and second is the percent of percent")
	}

	percent, _ := strconv.ParseFloat(os.Args[1], 64)
	ofPercent, _ := strconv.ParseFloat(os.Args[2], 64)
	fmt.Println("answer is: ", calc(percent, ofPercent))
}
