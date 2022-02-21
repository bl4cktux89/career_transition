package main

import (
	"fmt"
)

func main() {
	tabuada := 7
	for multiplicador := 1; multiplicador <= 10; multiplicador++ {
		resultado := tabuada * multiplicador
		fmt.Printf("%v x 2 = %v\n", multiplicador, resultado)
	}
}
