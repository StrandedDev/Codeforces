package main

import (
	"fmt"
)

func main() {
	var n, y, m, d, r_days int

	fmt.Scanf("%d", &n)

	y = n / 365
	r_days = n - (y * 365)

	m = r_days / 30
	d = r_days - (m * 30)
	fmt.Printf("%d years\n%d months\n%d days", y, m, d)

}
