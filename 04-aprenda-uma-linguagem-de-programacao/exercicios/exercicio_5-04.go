package main

import (
	"fmt"
)

func main() {
	x := struct {
		nome      string
		sabor     string
		ondeTem   []string
		vaiBemCom map[string]string
	}{
		nome:    "biscoito",
		sabor:   "doce",
		ondeTem: []string{"São Paulo", "Minas Gerais"},
		vaiBemCom: map[string]string{
			"no café da manhã": "leite com café",
			"no almoço":        "água",
			"na janta":         "chá de camomila",
		},
	}

	fmt.Println(x)

}
