package main

import (
	"fmt"
)

func main() {
	guess := 0
	switch {
	case guess == 0:
		fmt.Println("tá doido")
	case guess == 1:
		fmt.Println("tá maluco")
	case guess == 2:
		fmt.Println("tá fora da casinha")
	}
}
