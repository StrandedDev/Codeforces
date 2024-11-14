package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64

	fmt.Scanf("%f %f %f", &a, &b, &c)

	min := math.Min(math.Min(a, b), c)
	max := math.Max(math.Max(a, b), c)
	median := a + b + c - min - max

	fmt.Printf("%v\n%v\n%v\n\n%v\n%v\n%v", min, median, max, a, b, c)

}
