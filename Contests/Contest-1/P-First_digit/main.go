package main

import (
	"fmt"
)

func main() {
	var a uint32

	fmt.Scanf("%c", &a)

	if a%2 == 0 {
		fmt.Print("EVEN")
	} else if a%2 != 0 {
		fmt.Print("ODD")
	} else {
		fmt.Print("ERROR")
	}
}
