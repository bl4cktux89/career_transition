package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		z := i % 4
		fmt.Printf("%v : 4 resta %v\n", i, z)
	}
}
