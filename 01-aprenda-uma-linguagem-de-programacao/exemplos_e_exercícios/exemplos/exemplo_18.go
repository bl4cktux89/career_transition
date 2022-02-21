package main

import (
	"fmt"
)

func main() {
	for hours := 0; hours <= 12; hours++ {
		for minutes := 0; minutes < 60; minutes++ {
			fmt.Printf("%v:%v\n", hours, minutes)
		}
	}
}
