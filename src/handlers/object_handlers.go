package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Object struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetObjects(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetObjects")

		res, err := db.Query("SELECT id, name, description FROM object")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var objects []Object
		for res.Next() {
			var obj Object
			err = res.Scan(&obj.ID, &obj.Name, &obj.Description)
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			objects = append(objects, obj)
		}
		json.NewEncoder(w).Encode(objects)
	}
}

func GetObject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var obj Object
		err := db.QueryRow("SELECT id, name, description FROM object WHERE id=$1", id).Scan(&obj.ID, &obj.Name, &obj.Description)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(obj)
	}
}

func CreateObject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var obj Object
		err := json.NewDecoder(r.Body).Decode(&obj)
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// validate input
		if obj.Name == "" || obj.Description == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("INSERT INTO objects (name, description) VALUES ($1, $2)", obj.Name, obj.Description)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// func DeleteObject(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		id := vars["id"]

// 		_, err := db.Exec("DELETE FROM objects WHERE id=$1", id)
// 		if err != nil {
// 			log.Println(err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		w.WriteHeader(http.StatusNoContent)
// 	}
// }
