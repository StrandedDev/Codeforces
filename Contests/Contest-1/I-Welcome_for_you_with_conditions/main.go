package main

import (
	"fmt"
)

func main() {
	var a, b uint8

	fmt.Scanf("%d %d", &a, &b)

	if a >= b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
