package main

import (
	"fmt"
)

func main() {
	ss := [][]string{
		{
			"Bradock",
			"Silva",
			"Comer papel",
		},
		{
			"Sushi",
			"Silva",
			"Dormir",
		},
		{
			"Lana",
			"Sanesi",
			"Latir pra todo mundo",
		},
	}

	for _, valor := range ss {
		fmt.Println(valor[0])
		for _, item := range valor {
			fmt.Println("\t", item)
		}
	}
}
