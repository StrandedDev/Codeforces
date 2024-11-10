package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)

	fmt.Println(m%10 + n%10)
}
