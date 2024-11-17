package main

import (
	"fmt"
)

func main() {
	var f1, l1, f2, l2 string

	fmt.Scanf("%s %s\n%s %s", &f1, &l1, &f2, &l2)

	if l1 == l2 {
		fmt.Println("ARE Brothers")

	} else {
		fmt.Print("NOT")
	}
}
