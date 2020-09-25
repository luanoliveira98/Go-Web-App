package controllers

import (
	"net/http"
	"text/template"

	"../models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index é responsável por carregar a página inicial de protudos
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

// New é responsável por carregar a página de cadastro de produtos
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
