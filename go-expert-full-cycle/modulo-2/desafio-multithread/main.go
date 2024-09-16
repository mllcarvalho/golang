package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BrasilAPIResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Uf         string `json:"state"`
}

type ViaCEPResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func fetchFromBrasilAPI(cep string, ch chan<- string) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("BrasilAPI: Erro - %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- fmt.Sprintf("BrasilAPI: Erro - StatusCode %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("BrasilAPI: Erro ao ler resposta - %v", err)
		return
	}

	var data BrasilAPIResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		ch <- fmt.Sprintf("BrasilAPI: Erro ao decodificar JSON - %v", err)
		return
	}

	result := fmt.Sprintf("BrasilAPI: \nCEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nUF: %s",
		data.Cep, data.Logradouro, data.Bairro, data.Cidade, data.Uf)
	ch <- result
}

func fetchFromViaCEP(cep string, ch chan<- string) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("ViaCEP: Erro - %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- fmt.Sprintf("ViaCEP: Erro - StatusCode %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("ViaCEP: Erro ao ler resposta - %v", err)
		return
	}

	var data ViaCEPResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		ch <- fmt.Sprintf("ViaCEP: Erro ao decodificar JSON - %v", err)
		return
	}

	result := fmt.Sprintf("ViaCEP: \nCEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s",
		data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf)
	ch <- result
}

func main() {
	var cep string

	fmt.Print("Digite o CEP: ")
	fmt.Scan(&cep)

	resultCh := make(chan string)

	go fetchFromBrasilAPI(cep, resultCh)
	go fetchFromViaCEP(cep, resultCh)

	select {
	case result := <-resultCh:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: Timeout de 1 segundo atingido.")
	}
}
