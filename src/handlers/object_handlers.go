package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"src/util"

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
		util.HandleHttpError(w, err, util.InternalServerError, http.StatusInternalServerError)

		var objects []Object
		for res.Next() {
			var obj Object

			err = res.Scan(&obj.ID, &obj.Name, &obj.Description)
			util.HandleError(err, util.InternalServerError, http.StatusInternalServerError)

			objects = append(objects, obj)
		}
		json.NewEncoder(w).Encode(objects)
	}
}

func GetObject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetObject")
		vars := mux.Vars(r)
		id := vars["id"]

		var obj Object
		err := db.QueryRow("SELECT id, name, description FROM object WHERE id=$1", id).Scan(&obj.ID, &obj.Name, &obj.Description)
		util.HandleHttpError(w, err, util.InternalServerError, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(obj)
	}
}

func CreateObject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("CreateObject")
		var obj Object
		err := json.NewDecoder(r.Body).Decode(&obj)
		util.HandleError(err, util.BadRequest, http.StatusBadRequest)

		// validate input
		if obj.Name == "" || obj.Description == "" {
			http.Error(w, util.BadRequest, http.StatusBadRequest)
			return
		}

		_, err = db.Exec("INSERT INTO object (name, description) VALUES ($1, $2)", obj.Name, obj.Description)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func DeleteObject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("DeleteObject")
		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("DELETE FROM object WHERE id=$1", id)
		util.HandleError(err, util.InternalServerError, http.StatusInternalServerError)

		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateObject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("UpdateObject")
		vars := mux.Vars(r)
		id := vars["id"]

		var obj Object
		err := json.NewDecoder(r.Body).Decode(&obj)
		util.HandleError(err, util.BadRequest, http.StatusBadRequest)

		// validate input
		if obj.Name == "" || obj.Description == "" {
			http.Error(w, util.BadRequest, http.StatusBadRequest)
			return
		}

		_, err = db.Exec("UPDATE object SET name=$1, description=$2 WHERE id=$3", obj.Name, obj.Description, id)
		util.HandleError(err, util.InternalServerError, http.StatusInternalServerError)

		w.WriteHeader(http.StatusNoContent)
	}
}
