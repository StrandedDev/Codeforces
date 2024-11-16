package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var s, q string

	fmt.Scanf("%d %s %d %s %d", &a, &s, &b, &q, &c)

	if s == "+" {
		if a+b == c {
			fmt.Println("Yes")
		} else {
			fmt.Print(a + b)
		}
	} else if s == "*" {
		if a*b == c {
			fmt.Println("Yes")
		} else {
			fmt.Print(a * b)
		}
	} else if s == "-" {
		if a-b == c {
			fmt.Println("Yes")
		} else {
			fmt.Print(a - b)
		}
	} else {
		fmt.Println("Invalid input")
	}

}
