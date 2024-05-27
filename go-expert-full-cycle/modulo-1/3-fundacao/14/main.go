package main

import "fmt"

type Client struct {
	name string
}

type Conta struct {
	saldo float64
}

func (c Client) Andou() {
	c.name = "João Carvalho"
	fmt.Println(c.name, "andou")
}

func (c *Conta) Sacar(valor float64) float64 {
	c.saldo -= valor
	return c.saldo
}

func (c Conta) Simular(valor float64) float64 {
	c.saldo += valor
	fmt.Printf("Simulando: %.2f\n", c.saldo)
	return c.saldo
}

func NewConta() *Conta {
	return &Conta{}
}

func main() {

	cliente := Client{"João"}
	conta := Conta{200.0}

	cliente.Andou()
	fmt.Printf("Cliente: %s\n", cliente.name)

	conta.Sacar(50.0)
	fmt.Printf("Saldo: %.2f\n", conta.saldo)

	conta.Simular(500.00)
	fmt.Printf("Saldo: %.2f\n", conta.saldo)

	conta2 := NewConta()
	conta2.saldo = 1000.0

	fmt.Printf("Saldo conta 2: %.2f\n", conta2.saldo)

}
