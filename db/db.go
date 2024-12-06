package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// ConectaComBancoDeDados retorna uma conexão ativa com o banco de dados PostgreSQL
func ConectaComBancoDeDados() *sql.DB {
	dsn := "user=root dbname=go_crud password=1234 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}

	// Testa a conexão com o banco de dados
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	return db
}
