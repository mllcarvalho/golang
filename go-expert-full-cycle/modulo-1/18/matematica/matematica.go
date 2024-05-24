package matematica

type Number interface {
	~int | ~float64
}

func Soma[T Number](a T, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() {
	println("Andando...")
}
