package main

import (
	"fmt"
	"strings"
)

func main() {
	var x int32

	fmt.Scanf("%c", &x)
	v := int(x) // get ascii representation of x input

	if v >= 65 && v <= 90 { // input is uppercase
		fmt.Printf(strings.ToLower(string(x)))
	} else if v >= 97 && v <= 122 { // input is lowercase
		fmt.Printf("%s", strings.ToUpper(string(x)))
	}

}
