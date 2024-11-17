package main

import (
	"fmt"
)

func main() {
	var x float32

	fmt.Scanf("%f", &x)

	if x >= 0 && x <= 25 {
		fmt.Println("Interval [0,25]")
	} else if x > 25 && x <= 50 {
		fmt.Println("Interval (25,50]")
	} else if x > 50 && x <= 75 {
		fmt.Println("Interval (50,75]")
	} else if x > 75 && x <= 100 {
		fmt.Println("Interval (75,100]")
	} else {
		fmt.Println("Out of Intervals")
	}
}
