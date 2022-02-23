package main

import (
	"fmt"
)

func main() {
	tabuada := 6
	for multiplicador := 1; multiplicador <= 10; multiplicador++ {
		resultado := tabuada * multiplicador
		fmt.Printf("%v x %v = %v\n", tabuada, multiplicador, resultado)
	}
}
