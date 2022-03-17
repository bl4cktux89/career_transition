package main

import (
	"fmt"
)

func main() {
	born := 1989
	today := 2022
	for born <= today {
		fmt.Printf("%v\n", born)
		born++
	}
}
