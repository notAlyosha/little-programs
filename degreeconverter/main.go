package main

import "fmt"

func toFahrenheit(arg float64) float64 {
	return (arg * 1.8) + 32
}

func toCelsium(arg float64) float64 {
	return (arg - 32) * 0.55555555555
}

func main() {
	fmt.Printf("%.3f\n", toFahrenheit(43))
}
