package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64

	fmt.Scanln(&a, &b, &c)

	maxVal := math.Max(math.Max(a, b), c)
	minVal := math.Min(math.Min(a, b), c)

	fmt.Printf("%v %v", minVal, maxVal)

}
