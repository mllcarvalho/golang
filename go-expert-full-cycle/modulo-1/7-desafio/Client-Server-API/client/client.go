package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serverURL     = "http://localhost:8080/cotacao"
	clientTimeout = 1 * time.Millisecond
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
	defer cancel()

	bid, err := fetchCotacao(ctx)
	if err != nil {
		log.Println("Error fetching cotacao:", err)
		return
	}

	if _, err := saveCotacaoToFile(bid); err != nil {
		log.Println("Error saving cotacao to file:", err)
		return
	}

	log.Println("Cotacao saved to cotacao.txt")
}

func fetchCotacao(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result["bid"], nil
}

func saveCotacaoToFile(bid string) (len int, err error) {
	fileContent := fmt.Sprintf("Dólar: %s", bid)
	f, err := os.Create("cotacao.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write([]byte(fileContent))
}
