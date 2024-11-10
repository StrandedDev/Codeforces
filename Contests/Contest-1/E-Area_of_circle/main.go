package main

import (
	"fmt"
)

func main() {
	var r float64
	const pi = 3.141592653

	fmt.Scanf("%f", &r)

	fmt.Printf("%.9f", pi*(r*r))
}
