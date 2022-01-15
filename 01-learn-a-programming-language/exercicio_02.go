package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}
