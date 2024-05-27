package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo\n"))
	// tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	} else {
		println(tamanho)
	}

	f.Close()

	//leitura
	arquivo, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if n == 0 || err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}
