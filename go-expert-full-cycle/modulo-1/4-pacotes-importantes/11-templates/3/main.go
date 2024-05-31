package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Curso de Go", 20},
		{"Curso de Ruby", 40},
		{"Curso de Python", 60},
	})
	if err != nil {
		panic(err)
	}
}
