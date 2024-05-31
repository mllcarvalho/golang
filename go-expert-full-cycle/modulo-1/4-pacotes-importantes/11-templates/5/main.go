package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"Curso de Go", 20},
			{"Curso de Ruby", 40},
			{"Curso de Python", 60},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8000", nil)
}
