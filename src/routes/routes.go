package routes

import (
	"database/sql"
	"src/handlers"

	"github.com/gorilla/mux"
)

const (
	ObjectsRootPath = "/objects"
	ObjectIdPath    = "/objects/{id}"
)

func InitRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc(ObjectsRootPath, handlers.GetObjects(db)).Methods("GET")
	router.HandleFunc(ObjectIdPath, handlers.GetObject(db)).Methods("GET")
	router.HandleFunc(ObjectsRootPath, handlers.CreateObject(db)).Methods("POST")
	router.HandleFunc(ObjectIdPath, handlers.DeleteObject(db)).Methods("DELETE")
	router.HandleFunc(ObjectIdPath, handlers.UpdateObject(db)).Methods("PATCH")
}
