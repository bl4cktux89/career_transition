package main

import (
	"fmt"
)

func main() {
	slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for key, value := range slice {
		fmt.Println(key, value)
	}

	fmt.Printf("%T", slice)

}
