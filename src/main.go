package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"src/routes"
	"strconv"

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

	router := mux.NewRouter()
	routes.ObjectRoutes(router)

	apiPort := os.Getenv("API_PORT")

	port, port_err := strconv.Atoi(apiPort)
	if port_err != nil {
		log.Fatal("Error parsing port")
	}

	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
