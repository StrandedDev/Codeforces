package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B float64
	fmt.Scanf("%f %f", &A, &B)

	fmt.Printf("floor %v / %v = %v\n", A, B, math.Floor(A/B))
	fmt.Printf("ceil %v / %v = %v\n", A, B, math.Ceil(A/B))
	fmt.Printf("round %v / %v = %v", A, B, math.Round(A/B))
}
