package main

import "fmt"

func main() {

	fmt.Println("Primeira linha")
	defer fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
	defer fmt.Println("Quarta linha")
	fmt.Println("Quinta linha")
	fmt.Println("Sexta linha")
	defer fmt.Println("Setima linha")
	fmt.Println("Oitava linha")

}
