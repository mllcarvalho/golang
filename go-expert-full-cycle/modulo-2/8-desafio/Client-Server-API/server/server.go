package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	apiURL     = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	dbName     = "cotacoes.db"
	apiTimeout = 200 * time.Millisecond
	dbTimeout  = 10 * time.Millisecond
)

type Cotacao struct {
	USD struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, bid TEXT, timestamp DATETIME)")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cotacao", cotacaoHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fetchCotacao(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cotacao Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return "", err
	}

	return cotacao.USD.Bid, nil
}

func saveCotacao(ctx context.Context, db *sql.DB, bid string) error {
	query := "INSERT INTO cotacoes (bid, timestamp) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, bid, time.Now())
	return err
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), apiTimeout)
	defer cancel()

	bid, err := fetchCotacao(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch cotacao", http.StatusInternalServerError)
		log.Println("Error fetching cotacao:", err)
		return
	}

	dbCtx, dbCancel := context.WithTimeout(context.Background(), dbTimeout)
	defer dbCancel()

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	if err := saveCotacao(dbCtx, db, bid); err != nil {
		http.Error(w, "Failed to save cotacao", http.StatusInternalServerError)
		log.Println("Error saving cotacao:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": bid})
}
