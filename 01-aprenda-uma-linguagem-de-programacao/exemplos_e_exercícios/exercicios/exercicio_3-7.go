package main

import (
	"fmt"
)

func main() {
	guess := 2
	if guess == 0 {
		fmt.Println("tá doido")
	} else if guess == 1 {
		fmt.Println("tá maluco")
	} else {
		fmt.Println("tá fora da casinha")
	}
}
