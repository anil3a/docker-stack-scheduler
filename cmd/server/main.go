package main

import (
	"log"
	"net/http"
	"os"

	"stack-scheduler/internal/api"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Register API routes
	api.RegisterRoutes(mux)

	log.Println("Starting server on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
