package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "futebol"
	switch {
	case esporteFavorito == "futebol":
		fmt.Println("acertou")
	default:
		fmt.Println("errou")
	}
}
