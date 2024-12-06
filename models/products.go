package models

import (
	"database/sql"
	"log"

	"github.com/alura/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// BuscaTodosOsProdutos retorna todos os produtos do banco de dados
func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	rows, err := db.Query("SELECT id, nome, descricao, preco, quantidade FROM produtos")
	if err != nil {
		log.Fatalf("Erro ao buscar produtos: %v", err)
	}
	defer rows.Close()

	var produtos []Produto

	for rows.Next() {
		var produto Produto
		if err := rows.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade); err != nil {
			log.Fatalf("Erro ao ler produto: %v", err)
		}
		produtos = append(produtos, produto)
	}

	return produtos
}

// CriaNovoProduto insere um novo produto no banco de dados
func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	query := "INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)"
	if _, err := db.Exec(query, nome, descricao, preco, quantidade); err != nil {
		log.Fatalf("Erro ao criar produto: %v", err)
	}
}

// DeletaProduto remove um produto do banco de dados pelo ID
func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	query := "DELETE FROM produtos WHERE id = $1"
	if _, err := db.Exec(query, id); err != nil {
		log.Fatalf("Erro ao deletar produto: %v", err)
	}
}

// EditaProduto retorna um produto específico para edição
func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	query := "SELECT id, nome, descricao, preco, quantidade FROM produtos WHERE id = $1"
	row := db.QueryRow(query, id)

	var produto Produto
	if err := row.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade); err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("Produto com ID %s não encontrado", id)
		}
		log.Fatalf("Erro ao buscar produto: %v", err)
	}

	return produto
}

// AtualizaProduto atualiza os dados de um produto no banco de dados
func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	query := "UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5"
	if _, err := db.Exec(query, nome, descricao, preco, quantidade, id); err != nil {
		log.Fatalf("Erro ao atualizar produto: %v", err)
	}
}
