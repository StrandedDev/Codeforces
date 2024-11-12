package main

import (
	"fmt"
)

func main() {
	var a, s, b int32

	fmt.Scanf("%d%c%d", &a, &s, &b)

	operator := int(s)

	if operator == 43 {
		fmt.Print(a + b)
	} else if operator == 45 {
		fmt.Print(a - b)
	} else if operator == 42 {
		fmt.Print(a * b)
	} else if operator == 47 {
		fmt.Printf("%d", a/b)
	} else {
		fmt.Println("Invalid operator")
	}

}
