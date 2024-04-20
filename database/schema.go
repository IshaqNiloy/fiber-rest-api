package database

import "log"

// creates product table

func CreateProductTable() {
	_, err := DB.Query(`CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	amount integer,
	name text UNIQUE,
	description text,
	category text NOT NULL    
)
`)
	if err != nil {
		log.Fatal("error occurred while creating products table")
	}
}
