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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
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
