package main

import (
	"net/http"

	"go_crud/routes"
)

func main() {
	// Configura as rotas da aplicação
	routes.CarregaRotas()

	// Inicia o servidor na porta 8000
	http.ListenAndServe(":8000", nil)
}
