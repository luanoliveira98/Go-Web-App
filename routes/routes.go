package routes

import (
	"net/http"

	"../controllers"
)

// CarregaRotas é resposável por carregar todoas as rotas do sistema
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
