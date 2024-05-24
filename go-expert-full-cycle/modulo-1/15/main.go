package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = "Matheus"

	ShowType(x)
	ShowType(y)

}

func ShowType(t interface{}) {
	fmt.Printf("O tipo é: %T e o valor %v\n", t, t)
}
