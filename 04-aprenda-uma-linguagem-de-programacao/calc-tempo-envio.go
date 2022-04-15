package main

import (
	"fmt"
	"math"
)

var fileQuantity int = 350
var fileSize int = 100
var linkSpeed = 1

func main() {
	pow1 := math.Pow(2, 10)
	pow2 := math.Pow(10, 6)
	tranferTime := (fileQuantity * fileSize * int(pow1) * 8) / (linkSpeed * int(pow2))
	fmt.Println(tranferTime)
}
