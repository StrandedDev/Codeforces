package main

import (
	"fmt"
)

func main() {
	var x, y float64

	fmt.Scanf("%f %f", &x, &y)

	if x > 0 && y > 0 {
		fmt.Print("Q1")
	} else if x < 0 && y > 0 {
		fmt.Print("Q2")
	} else if x < 0 && y < 0 {
		fmt.Print("Q3")
	} else if x > 0 && y < 0 {
		fmt.Print("Q4")
	} else if x == 0 && y == 0 {
		fmt.Print("Origem")
	} else if x == 0 && (y > 0 || y < 0) {
		fmt.Print("Eixo Y")
	} else if y == 0 && (x > 0 || x < 0) {
		fmt.Print("Eixo X")
	} else {
		fmt.Print("Error")
	}
}
