package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"go_crud/models"
)

var templates = template.Must(template.ParseGlob("views/*.html"))

// Index renderiza a página inicial com todos os produtos
func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	if err := templates.ExecuteTemplate(w, "Index", allProducts); err != nil {
		log.Printf("Error rendering Index template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// New renderiza o formulário para criar um novo produto
func New(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "New", nil); err != nil {
		log.Printf("Error rendering New template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Insert processa os dados enviados pelo formulário e cria um novo produto
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Printf("Error converting price: %v", err)
			http.Error(w, "Invalid price format. Please use a dot (.) as the decimal separator.", http.StatusBadRequest)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Printf("Error converting quantity: %v", err)
			http.Error(w, "Invalid quantity format. Please provide a valid number.", http.StatusBadRequest)
			return
		}

		models.CreateNewProduct(name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Delete remove um produto com base no ID fornecido
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Product ID not provided", http.StatusBadRequest)
		return
	}

	models.DeleteProduct(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Edit renderiza o formulário para editar um produto existente
func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Product ID not provided", http.StatusBadRequest)
		return
	}

	product := models.EditProduct(id)
	if err := templates.ExecuteTemplate(w, "Edit", product); err != nil {
		log.Printf("Error rendering Edit template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Update processa os dados enviados pelo formulário e atualiza um produto existente
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error converting ID: %v", err)
			http.Error(w, "Invalid ID format. Please provide a valid number.", http.StatusBadRequest)
			return
		}

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Printf("Error converting price: %v", err)
			http.Error(w, "Invalid price format. Please use a dot (.) as the decimal separator.", http.StatusBadRequest)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Printf("Error converting quantity: %v", err)
			http.Error(w, "Invalid quantity format. Please provide a valid number.", http.StatusBadRequest)
			return
		}

		models.UpdateProduct(idInt, name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
