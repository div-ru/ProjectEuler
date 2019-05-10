package main

import (
	"fmt"
)

func main() {

	var item1 = 1
	var item2 = 2
	var tmp, sum int

	for item2 <= 4000000 {
		if item2%2 == 0 {
			sum += item2
		}
		tmp = item1 + item2
		item1 = item2
		item2 = tmp
	}

	fmt.Println("Sum:", sum)
}
