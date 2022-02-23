package main

import (
	"fmt"
)

func main() {
	born := 1989
	today := 2022

	for {
		if born > today {
			break
		}
		fmt.Println(born)
		born++
	}
}
