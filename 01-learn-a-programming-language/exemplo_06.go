package main

import (
	"fmt"
)

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Printf("%v\n%T\n", x, x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
