package main

import (
	"fmt"
)

func main() {
	slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2 : len(slice)-1])

}
