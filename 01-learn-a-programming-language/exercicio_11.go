package main

import (
	"fmt"
)

func main() {
	s := "hello world"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println(" ")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

}
