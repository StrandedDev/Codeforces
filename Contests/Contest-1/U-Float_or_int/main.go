package main

import (
	"fmt"
)

func main() {
	var x int
	var y float32

	fmt.Scanf("%d", &x)
	fmt.Scanf("%f", &y)

	if y == 0 {
		fmt.Printf("int %d", x)
	} else {
		fmt.Printf("float %d 0.%v", x, y)
	}

}
