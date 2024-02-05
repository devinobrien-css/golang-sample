package util

import (
	"log"
	"net/http"
)

const (
	InternalServerError = "Internal Server Error"
	BadRequest          = "Bad Request"
)

func HandleHttpError(writer http.ResponseWriter, err error, msg string, status int) {
	if err != nil {
		log.Println(err)
		http.Error(writer, msg, status)
		return
	}
}

func HandleError(err error, msg string, status int) {
	if err != nil {
		log.Println(err)
	}
}
