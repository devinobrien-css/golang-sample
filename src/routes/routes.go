package routes

import (
	"database/sql"
	"src/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/objects", handlers.GetObjects(db)).Methods("GET")
	router.HandleFunc("/objects/{id}", handlers.GetObject(db)).Methods("GET")
	// router.HandleFunc("/objects", CreateObject(db)).Methods("POST")
	// router.HandleFunc("/objects/{id}", DeleteObject(db)).Methods("DELETE")
}
