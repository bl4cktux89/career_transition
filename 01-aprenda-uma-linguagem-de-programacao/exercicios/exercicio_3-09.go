package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "volei"
	switch esporteFavorito {
	case "futebol":
		fmt.Println("acertou")
	default:
		fmt.Println("errou")
	}
}
