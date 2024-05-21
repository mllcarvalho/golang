package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Endereco
}

func (c *Client) DesativarCliente() {
	c.Active = false
	fmt.Println("Cliente desativado")
}

func main() {

	matheus := Client{
		Name:   "Matheus",
		Age:    20,
		Active: true,
		Endereco: Endereco{
			Logradouro: "Rua dos Bobos",
			Numero:     0,
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}

	matheus.DesativarCliente()

	fmt.Println(matheus.Active)

}
