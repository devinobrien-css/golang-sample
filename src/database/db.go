package database

import (
	"database/sql"
	"net/http"
	"src/util"
)

func InitDB(dsn string) *sql.DB {
	// Open a connection to the database
	conn, err := sql.Open("postgres", dsn)
	util.HandleError(err, "Failed to open a connection to the database", http.StatusInternalServerError)

	// Test the connection
	err = conn.Ping()
	util.HandleError(err, "Failed to ping the database", http.StatusInternalServerError)

	return conn
}
