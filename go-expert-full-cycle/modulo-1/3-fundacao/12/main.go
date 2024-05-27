package main

import "fmt"

func main() {

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := a

	c := 50
	var d *int = &c
	d = &c
	*d = 100

	fmt.Println(a, b)
	fmt.Println(c)

}
