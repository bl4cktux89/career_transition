package main

import (
	"fmt"
)

// declaração do tipo
type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	// declaração mais detalhada
	cliente1 := cliente{
		nome:      "John",
		sobrenome: "Textor",
		fumante:   true,
	}

	// declaração mais concisa
	cliente2 := cliente{"Joana", "Darc", false}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
