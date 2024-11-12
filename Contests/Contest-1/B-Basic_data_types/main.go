package main

import (
	"fmt"
)

func main() {

	var a int
	var b int
	var c string
	var d float32
	var e float64

	fmt.Scanf("%d %d %s %f %f", &a, &b, &c, &d, &e)
	fmt.Printf("%d\n%d\n%s\n%.2f\n%.1f\n", a, b, c, d, e)

}
