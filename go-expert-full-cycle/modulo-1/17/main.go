package main

import "fmt"

type Number interface {
	~int | ~float64
}

type MyNumber int

func main() {

	m1 := map[string]int{
		"Matheus": 10,
		"João":    20,
		"Maria":   30,
	}

	m2 := map[string]float64{
		"Matheus": 10.5,
		"João":    20.5,
		"Maria":   30.5,
	}

	m3 := map[string]MyNumber{
		"Matheus": 10,
		"João":    20,
		"Maria":   30,
	}

	println(Soma(m1))
	fmt.Println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10.5))

}

func Soma[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func Compara[T Number](a T, b T) bool {
	return a == b
}
