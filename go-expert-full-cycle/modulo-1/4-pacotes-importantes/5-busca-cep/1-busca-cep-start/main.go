package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", cep, err)
			os.Exit(1)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: parsing %s: %v\n", cep, err)
		}
		fmt.Println(data)

		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf))

		fmt.Println("Arquivo cidade.txt criado com sucesso!")
		fmt.Println("Cidade: ", data.Localidade)

	}

}
