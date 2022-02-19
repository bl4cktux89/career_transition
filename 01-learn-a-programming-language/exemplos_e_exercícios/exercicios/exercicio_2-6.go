package main

import (
	"fmt"
)

const (
	_ = iota + 2022
	a
	b
	c
	d
)

func main() {
	fmt.Printf("%v\t%v\t%v\t%v\n", a, b, c, d)
}
