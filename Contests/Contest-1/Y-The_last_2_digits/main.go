package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, res int64

	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// a2 := a % 100
	// b2 := b % 100
	// c2 := c % 100
	// d2 := d % 100

	res = ((a % 100) * (b % 100) * (c % 100) * (d % 100)) % 100

	fmt.Printf("%02d", res)

}
