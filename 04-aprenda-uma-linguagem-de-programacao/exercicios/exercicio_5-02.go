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
	mymap := make(map[string]pessoa)
	mymap["Silva"] = pessoa{
		nome:      "Vinícius",
		sobrenome: "Silva",
		sabores:   []string{"brigadeiro", "flocos", "morango"},
	}
	for _, valor := range mymap {
		fmt.Println("meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabores {
			fmt.Println(" * ", valor)
		}
	}
}
