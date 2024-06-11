package main

import (
	"fmt"

	"github.com/mllcarvalho/golang/go-expert-full-cycle/modulo-2/9-packaging/3/math"
)

func main() {
	fmt.Println("Hello, World!")

	result := math.Math{A: 10, B: 20}.Sum()
	fmt.Println(result)
}
