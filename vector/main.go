// ========================================================================
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// ========================================================================
type Point struct {
	X float64
	Y float64
}

func calc(first, second *Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(second.X-first.X), 2) + math.Pow(math.Abs(second.Y-first.Y), 2))
}

// ========================================================================
func main() {

	if len(os.Args) < 5 {
		panic("Format is x1 y1 x2 y2")
	}

	p1 := new(Point)
	p2 := new(Point)

	p1.X, _ = strconv.ParseFloat(os.Args[1], 64)
	p1.Y, _ = strconv.ParseFloat(os.Args[2], 64)
	p2.X, _ = strconv.ParseFloat(os.Args[3], 64)
	p2.Y, _ = strconv.ParseFloat(os.Args[4], 64)

	fmt.Println(calc(p1, p2))
}
