package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// product := NewProduct("Notebook", 1899.90)
	// err = insertProduct(db, product)
	// if err != nil {
	// 	panic(err)
	// }
	// product.Price = 100.00
	// err = updateProduct(db, product)
	// if err != nil {
	// 	panic(err)
	// }

	// p, err := selectOneProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: %+v\n", p)

	err = deleteProduct(db, "3a2f8b60-fe2a-41b5-ad5f-e793503250cf")
	if err != nil {
		panic(err)
	}

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %+v\n", p)
	}
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		panic(err)
	}
	return nil
}

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	var p Product
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	var products []Product
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return nil
}
