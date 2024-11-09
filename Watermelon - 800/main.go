package main

import (
	"fmt"
)

func getValue(w uint16) string {
	if w%2 == 0 && w > 2 {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	var w uint16
	fmt.Scanln(&w)
	fmt.Println(getValue(w))
}
