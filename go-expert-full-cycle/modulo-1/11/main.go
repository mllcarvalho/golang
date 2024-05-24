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

func (c *Client) DesativarCliente() string {
	c.Active = false
	return "Cliente desativado"
}

func main() {

	matheus := Client{
		Name:   "Matheus",
		Age:    20,
		Active: true,
		Endereco: Endereco{
			Logradouro: "Rua dos Bobos",
			Numero:     50,
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}

	retorno := matheus.DesativarCliente()

	fmt.Println(retorno)
	fmt.Println(matheus.Active)

}
