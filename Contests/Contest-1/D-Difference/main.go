package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int64
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Println("Difference =", ((a * b) - (c * d)))
}
