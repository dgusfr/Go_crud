package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/alura/models"
)

var templates = template.Must(template.ParseGlob("views/*.html"))

// Index exibe a lista de produtos
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	if err := templates.ExecuteTemplate(w, "Index", produtos); err != nil {
		log.Printf("Erro ao renderizar template Index: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}

// New exibe o formulário para criar um novo produto
func New(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "New", nil); err != nil {
		log.Printf("Erro ao renderizar template New: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}

// Insert insere um novo produto no banco de dados
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Printf("Erro ao converter preço: %v", err)
			http.Error(w, "Preço inválido", http.StatusBadRequest)
			return
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Printf("Erro ao converter quantidade: %v", err)
			http.Error(w, "Quantidade inválida", http.StatusBadRequest)
			return
		}

		models.CriaNovoProduto(nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Delete remove um produto pelo ID
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID do produto não fornecido", http.StatusBadRequest)
		return
	}

	models.DeletaProduto(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Edit exibe o formulário de edição de um produto
func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID do produto não fornecido", http.StatusBadRequest)
		return
	}

	produto := models.EditaProduto(id)
	if err := templates.ExecuteTemplate(w, "Edit", produto); err != nil {
		log.Printf("Erro ao renderizar template Edit: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}

// Update atualiza os dados de um produto no banco de dados
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erro ao converter ID: %v", err)
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Printf("Erro ao converter preço: %v", err)
			http.Error(w, "Preço inválido", http.StatusBadRequest)
			return
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Printf("Erro ao converter quantidade: %v", err)
			http.Error(w, "Quantidade inválida", http.StatusBadRequest)
			return
		}

		models.AtualizaProduto(idInt, nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
