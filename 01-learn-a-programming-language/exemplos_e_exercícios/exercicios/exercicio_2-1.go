package main

import (
	"fmt"
)

func main() {
	x := 1406
	fmt.Printf("o valor de %v em binário é: %b\n", x, x)
	fmt.Printf("o valor de %v em decimal é: %d\n", x, x)
	fmt.Printf("o valor de %v em octal é: %o\n", x, x)
	fmt.Printf("o valor de %v em hexadecimal é: %#x\n", x, x)
}
