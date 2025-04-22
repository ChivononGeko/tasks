package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // драйвер PostgreSQL
)

func Connect(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
		return nil, err
	}

	return db, nil
}
