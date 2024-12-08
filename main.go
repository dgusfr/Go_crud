package main

import (
	"log"
	"net/http"
)

func main() {
	// Carrega as rotas da aplicação
	routes.CarregaRotas()

	// Define a porta e inicia o servidor
	port := ":8000"
	log.Printf("Servidor iniciado na porta %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
