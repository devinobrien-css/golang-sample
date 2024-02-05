package database

import (
	"database/sql"
	"log"
)

func InitDB(dsn string) *sql.DB {
	// Open a connection to the database
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to the database")
	}
	// defer conn.Close()
	log.Default().Println("Connected to database")

	// Test the connection
	err = conn.Ping()
	if err != nil {
		panic("Failed to ping the database")
	}

	return conn
}
