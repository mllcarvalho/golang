package main

import (
	"fmt"
)

func main() {

	valor := sum(1, 10, 30, 50, 60)

	fmt.Println(valor)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
