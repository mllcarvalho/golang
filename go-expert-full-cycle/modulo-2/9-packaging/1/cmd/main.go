package main

import (
	"fmt"

	"github.com/mllcarvalho/golang/go-expert-full-cycle/modulo-2/9-packaging/1/math"
)

func main() {
	fmt.Println("Hello World!")
	m := math.Math{A: 10, B: 20}
	fmt.Println(m.Sum())
}
