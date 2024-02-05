package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"src/database"
	"src/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()

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

	// fetch server port from .env file
	apiPort := os.Getenv("API_PORT")

	// start server
	fmt.Printf("Server is running on http://localhost:%s\n", apiPort)
	serverErr := http.ListenAndServe(fmt.Sprintf(":%s", apiPort), router)
	if serverErr != nil {
		fmt.Println("Error:", serverErr)
	}
}
