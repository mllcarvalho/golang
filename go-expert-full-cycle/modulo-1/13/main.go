package main

import "fmt"

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {

	a := 10
	b := 20

	fmt.Println(a)
	fmt.Println(soma(&a, &b))
	fmt.Println(a)

}
