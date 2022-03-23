package main

import (
	"fmt"
)

// declaração do tipo
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type profissao struct {
	pessoa
	titulo  string
	salario float64
}

func main() {
	// declaração mais detalhada
	pessoa1 := pessoa{
		nome:      "John",
		sobrenome: "Textor",
		idade:     78,
	}

	pessoa2 := profissao{
		pessoa: pessoa{
			nome:      "Vinícius",
			sobrenome: "Silva",
			idade:     32,
		},
		titulo:  "Analista de Redes",
		salario: 4125.00,
	}

	// declaração mais concisa
	pessoa3 := pessoa{"Joana", "Darc", 56}

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Valdir",
		idade: 56,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
	fmt.Println(x)
}
