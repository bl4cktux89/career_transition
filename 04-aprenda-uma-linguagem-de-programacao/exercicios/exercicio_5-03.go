package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	carraoDoTio := sedan{veiculo{4, "vermelho"}, true}
	carroDoTrampo := caminhonete{veiculo{4, "branca"}, true}
	fmt.Println(carraoDoTio)
	fmt.Println(carroDoTrampo)

}
