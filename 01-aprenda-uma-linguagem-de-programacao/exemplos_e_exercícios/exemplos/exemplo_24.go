package main

import (
	"fmt"
)

func main() {
	slice := []int{20, 21, 22, 23}

	total := 0
	for _, valor := range slice {
		total += valor
	}
	fmt.Println("o valor total é:", total)
}
