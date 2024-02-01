package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Object struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func GetObjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetObjects")
		// vars := mux.Vars(r)
		// id := vars["id"]

		// var obj Object
		// err := db.QueryRow("SELECT id, name, value FROM objects WHERE id=$1", id).Scan(&obj.ID, &obj.Name, &obj.Value)
		// if err != nil {
		// 	log.Println(err)
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	return
		// }
		var obj = Object{
			ID:    1,
			Name:  "asparagandasaurus",
			Value: "asparagandasaurus",
		}

		json.NewEncoder(w).Encode(obj)
	}
}

// func GetObject(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		id := vars["id"]

// 		var obj Object
// 		err := db.QueryRow("SELECT id, name, value FROM objects WHERE id=$1", id).Scan(&obj.ID, &obj.Name, &obj.Value)
// 		if err != nil {
// 			log.Println(err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		json.NewEncoder(w).Encode(obj)
// 	}
// }

// func CreateObject(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var obj Object
// 		err := json.NewDecoder(r.Body).Decode(&obj)
// 		if err != nil {
// 			log.Println(err)
// 			http.Error(w, "Bad Request", http.StatusBadRequest)
// 			return
// 		}

// 		_, err = db.Exec("INSERT INTO objects (name, value) VALUES ($1, $2)", obj.Name, obj.Value)
// 		if err != nil {
// 			log.Println(err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}

// 		w.WriteHeader(http.StatusCreated)
// 	}
// }

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
