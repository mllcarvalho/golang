package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", Blog{})
	http.ListenAndServe(":8080", mux)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type Blog struct{}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog"))
}
