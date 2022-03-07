package main

import (
	"fmt"
)

func main() {
	array := [5]int{9, 8, 7, 6, 5}

	for key, value := range array {
		fmt.Println(key, value)
	}

	fmt.Printf("%T\n", array)
}
