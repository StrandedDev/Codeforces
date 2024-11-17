package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64

	fmt.Scanf("%f %f %f %f", &a, &b, &c, &d)

	if b*math.Log(a) > d*math.Log(c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
