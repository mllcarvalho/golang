package main

func main() {

	a := 4
	b := 2
	c := 3

	if a < b {
		println("a é menor que b")
	} else {
		println("b é menor que a")
	}

	if a < b && b < c {
		println("a é menor que b e b é menor que c")
	}

	switch a {
	case 1:
		println("a é 1")
	case 2:
		println("a é 2")
	case 3:
		println("a é 3")
	default:
		println("a não é 1, 2 ou 3")
	}
}
