package main

import (
	"fmt"
	"net/http"
	"os"
	"src/database"
	"src/routes"
	"src/util"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	util.HandleError(err, "Error loading .env file", http.StatusInternalServerError)

	// establish database connection
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db := database.InitDB(dsn)

	// create server router
	router := mux.NewRouter()
	routes.InitRoutes(router, db)

	// start server
	apiPort := os.Getenv("API_PORT")
	fmt.Printf("Server is running on http://localhost:%s\n", apiPort)
	serverErr := http.ListenAndServe(fmt.Sprintf(":%s", apiPort), router)
	util.HandleError(serverErr, util.InternalServerError, http.StatusInternalServerError)
}
