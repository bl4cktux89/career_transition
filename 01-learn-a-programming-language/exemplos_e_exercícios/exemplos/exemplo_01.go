package main

import (
	"fmt"
)

type hotdog int

var b hotdog

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", b)
}
