package main

import (
	"fmt"
)

func main() {

	var item1 uint32 = 1
	var item2 uint32 = 2
	var tmp uint32
	var sum uint32

	for item2 <= uint32(4000000) {
		if item2%2 == 0 {
			sum += item2
		}
		tmp = item1 + item2
		item1 = item2
		item2 = tmp
	}

	fmt.Println("Sum:", sum)
}
