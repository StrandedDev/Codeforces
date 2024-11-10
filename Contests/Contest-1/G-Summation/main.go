package main

import (
	"fmt"
)

func main() {

	var N uint64
	fmt.Scanf("%d", &N)

	sum := (N * (N + 1)) / 2

	fmt.Println(sum)

}
