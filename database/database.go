package database

import (
	"database/sql"
	"fmt"
	"github.com/IshaqNiloy/go-rest-api/config"
	_ "github.com/lib/pq"
)

// Database instance

var DB *sql.DB

// Connect function
func Connect() error {
	var err error

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.Config("DB_USER"),
		config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"),
		config.Config("DB_NAME"), config.Config("SSL_MODE"))

	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	CreateProductTable()
	fmt.Println("Connection Opened to Database")
	return nil
}
