package main

import (
	"html/template"
	"net/http"
)

// Produto padrão da loja
type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Azul, bem bonita", 39, 5},
		{"Tênis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
