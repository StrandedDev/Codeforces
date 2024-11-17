package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("%v * %v = %v\n", a, b, a*b)
	fmt.Printf("%v - %v = %v\n", a, b, a-b)
}
