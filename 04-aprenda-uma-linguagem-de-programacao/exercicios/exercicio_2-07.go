package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\t\n", x, x, x)
	y := 31337
	fmt.Printf("%d\t%x\t%#x\t\n", y, y, y)
}
