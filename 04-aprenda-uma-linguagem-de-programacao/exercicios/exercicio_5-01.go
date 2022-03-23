package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := pessoa{
		nome:      "Vinícius",
		sobrenome: "Silva",
		sabores:   []string{"brigadeiro", "flocos", "morango"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}
}
