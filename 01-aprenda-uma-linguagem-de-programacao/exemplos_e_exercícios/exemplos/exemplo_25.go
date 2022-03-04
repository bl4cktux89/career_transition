package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		{1, 2, 3, 4, 5, 6},
		{11, 12, 13, 14, 15, 16},
		{21, 22, 23, 24, 25, 26},
	}
	fmt.Println((ss[2][5]))
}
