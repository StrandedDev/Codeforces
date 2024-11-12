package main

import (
	"fmt"
)

func main() {
	var x int32

	fmt.Scanf("%c", &x)

	dec_value := int(x)

	if dec_value >= 48 && dec_value <= 57 {
		fmt.Println("IS DIGIT")
	} else if dec_value >= 65 && dec_value <= 90 {
		fmt.Printf("ALPHA\nIS CAPITAL")
	} else if dec_value >= 97 && dec_value <= 122 {
		fmt.Printf("ALPHA\nIS SMALL")
	} else {
		fmt.Println("NOT ALPHA OR DIGIT")
	}
}
