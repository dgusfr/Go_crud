package routes

import (
	"net/http"

	"github.com/Alura/controllers"
)

// CarregaRotas configura todas as rotas da aplicação
func CarregaRotas() {
	routes := map[string]http.HandlerFunc{
		"/":        controllers.Index,
		"/new":     controllers.New,
		"/insert":  controllers.Insert,
		"/delete":  controllers.Delete,
		"/edit":    controllers.Edit,
		"/update":  controllers.Update,
	}

	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}
}
