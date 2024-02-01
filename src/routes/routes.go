package routes

import (
	"fmt"
	"src/handlers"

	"github.com/gorilla/mux"
)

func foo() {
	fmt.Print("foo")
}

func ObjectRoutes(router *mux.Router) {
	router.HandleFunc("/objects", handlers.GetObjects()).Methods("GET")
	// router.HandleFunc("/objects/{id}", GetObject(db)).Methods("GET")
	// router.HandleFunc("/objects", CreateObject(db)).Methods("POST")
	// router.HandleFunc("/objects/{id}", DeleteObject(db)).Methods("DELETE")
}
