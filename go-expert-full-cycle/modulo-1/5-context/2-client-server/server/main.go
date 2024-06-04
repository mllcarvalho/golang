package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("handler started")
	defer log.Println("handler ended")

	select {
	case <-ctx.Done():
		log.Println("Request cancell")
	case <-time.After(5 * time.Second):
		log.Println("Request processed")
		w.Write([]byte("Request processed"))
	}
}
