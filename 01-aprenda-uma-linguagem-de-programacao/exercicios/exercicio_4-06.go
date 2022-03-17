package main

import (
	"fmt"
)

func main() {
	uf := make([]string, 26, 26)
	uf = []string{"AC", "AL", "AP", "AM", "BA", "CE", "ES",
		"GO", "MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE",
		"PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}
	fmt.Println(len(uf), cap(uf))
	for state := 0; state < len(uf); state++ {
		fmt.Println(uf[state])
	}
}
