package controllers

import (
	"log"
	"net/http"
	"strconv"
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

// Insert é responsável por criar um novo produto
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

// Delete é responsável por remover um produto
func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

// Edit é responsábel por carregar a página de edição de produto
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}
