package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Alura/models/products"
)

var templates = template.Must(template.ParseGlob("views/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := products.GetAllProducts()
	if err := templates.ExecuteTemplate(w, "Index", allProducts); err != nil {
		log.Printf("Error rendering Index template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "New", nil); err != nil {
		log.Printf("Error rendering New template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Printf("Error converting price: %v", err)
			http.Error(w, "Invalid price format", http.StatusBadRequest)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Printf("Error converting quantity: %v", err)
			http.Error(w, "Invalid quantity format", http.StatusBadRequest)
			return
		}

		products.CreateNewProduct(name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Product ID not provided", http.StatusBadRequest)
		return
	}

	products.DeleteProduct(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Product ID not provided", http.StatusBadRequest)
		return
	}

	product := products.EditProduct(id)
	if err := templates.ExecuteTemplate(w, "Edit", product); err != nil {
		log.Printf("Error rendering Edit template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

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
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Printf("Error converting price: %v", err)
			http.Error(w, "Invalid price format", http.StatusBadRequest)
			return
		}

		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Printf("Error converting quantity: %v", err)
			http.Error(w, "Invalid quantity format", http.StatusBadRequest)
			return
		}

		products.UpdateProduct(idInt, name, description, priceFloat, quantityInt)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
