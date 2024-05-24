package main

import (
	"fmt"

	"github.com/go-expert/matematica"
	"github.com/google/uuid"
)

func main() {

	s := matematica.Soma(10, 10)

	fmt.Printf("%v\n", s)

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Printf("%v\n", carro.Marca)

	fmt.Println(uuid.New().String())

	carro.Andar()
}
