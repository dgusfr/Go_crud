package routes

import (
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", productsController.Index)
	http.HandleFunc("/new", productsController.New)
	http.HandleFunc("/insert", productsController.Insert)
	http.HandleFunc("/delete", productsController.Delete)
	http.HandleFunc("/edit", productsController.Edit)
	http.HandleFunc("/update", productsController.Update)
}
