package products

import (
	"database/sql"
	"log"

	"github.com/Alura/database"
)

type Product struct {
	Id         int
	Name       string
	Description string
	Price      float64
	Quantity   int
}

func GetAllProducts() []Product {
	db := database.ConnectToDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, description, price, quantity FROM products")
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity); err != nil {
			log.Fatalf("Error reading product: %v", err)
		}
		products = append(products, product)
	}

	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := database.ConnectToDatabase()
	defer db.Close()

	query := "INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)"
	if _, err := db.Exec(query, name, description, price, quantity); err != nil {
		log.Fatalf("Error creating product: %v", err)
	}
}

func DeleteProduct(id string) {
	db := database.ConnectToDatabase()
	defer db.Close()

	query := "DELETE FROM products WHERE id = $1"
	if _, err := db.Exec(query, id); err != nil {
		log.Fatalf("Error deleting product: %v", err)
	}
}

func EditProduct(id string) Product {
	db := database.ConnectToDatabase()
	defer db.Close()

	query := "SELECT id, name, description, price, quantity FROM products WHERE id = $1"
	row := db.QueryRow(query, id)

	var product Product
	if err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity); err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("Product with ID %s not found", id)
		}
		log.Fatalf("Error fetching product: %v", err)
	}

	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := database.ConnectToDatabase()
	defer db.Close()

	query := "UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5"
	if _, err := db.Exec(query, name, description, price, quantity, id); err != nil {
		log.Fatalf("Error updating product: %v", err)
	}
}
