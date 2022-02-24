package main

import (
	"fmt"
)

func main() {
	slice := []string{"banana", "maçã", "jaca", "pêssego"}
	slice = append(slice, "melancia")

	for i, x := range slice {
		fmt.Println("no índice", i, "temos o valor:", x)
	}
}
