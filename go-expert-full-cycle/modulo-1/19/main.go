package main

func main() {

	for i := 0; i < 4; i++ {
		println(i)
	}

	numeros := []string{"zero", "um", "dois", "três"}

	for k, v := range numeros {
		println(k, v)
	}

	i := 0
	for i < 4 {
		println(i)
		i++
	}

	for {
		println("Loop infinito")
		break //sem o break, o loop seria infinito
	}
}
