package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver para MySQL
)

// ConnectToDatabase conecta ao banco de dados MySQL
func ConnectToDatabase() *sql.DB {
	// Substitua as credenciais com as suas configurações
	connStr := "root:1234@tcp(127.0.0.1:3306)/go_crud"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Error opening connection to database: %v", err)
	}

	// Verifica se a conexão foi bem-sucedida
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to the database!")
	return db
}
