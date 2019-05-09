package main

import (
	"fmt"
)

func main() {

	var sum uint32

	for i := 0; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += uint32(i)
		}
	}

	fmt.Println("Sum:", sum)
}
