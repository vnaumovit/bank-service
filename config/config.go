package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() *sql.DB {
	connStr := "user=postgres password=postgres dbname=bank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	InitAccountsTable(db)
	return db
}

func InitAccountsTable(db *sql.DB) {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS accounts (
        id SERIAL PRIMARY KEY,
        name TEXT UNIQUE NOT NULL,
        balance NUMERIC NOT NULL DEFAULT 0
    );
    `
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	fmt.Println("Table created or already exists")
}
