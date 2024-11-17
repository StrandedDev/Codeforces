package main

import (
	"fmt"
)

func main() {
	var x, y, a, b int

	fmt.Scanf("%d %d %d %d", &x, &y, &a, &b)

	if x <= a && a <= y && y <= b {
		fmt.Println(a, y)
	} else if x <= a && a <= b && b <= y {
		fmt.Println(a, b)
	} else if a <= x && x <= b && b <= y {
		fmt.Println(x, b)
	} else if a <= x && x <= y && y <= b {
		fmt.Println(x, y)
	} else {
		fmt.Println(-1)
	}

}
