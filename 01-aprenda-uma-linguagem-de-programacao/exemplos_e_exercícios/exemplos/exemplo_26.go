package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9995678,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 4444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])
}
