package main

import (
	"fmt"
)

func main() {
	mapa := map[string][]string{
		"silva_vinicius": {
			"lol", "devops",
		},
		"sanesi_carol": {
			"make", "docinho",
		},
		"silva_leia": {
			"destruição", "caos",
		},
	}
	delete(mapa, "silva_leia")
	mapa["sanesi_maria"] = []string{"trabalhar", "ir à missa"}
	for k, v := range mapa {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, "\t-\t", h)
		}
	}
}
