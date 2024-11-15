package main

import (
	"fmt"
)

func main() {

	var a, b int8
	var s string

	fmt.Scanf("%d %s %d", &a, &s, &b)

	if s == ">" {
		if a > b {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	} else if s == "<" {
		if a < b {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	} else if s == "=" {
		if a == b {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	} else {
		fmt.Println("Invalid input")
	}

}
