package main

import (
	"log"
	"net/http"

	"github.com/RPDJF/EnvMesh-backend/internal/routes"
)

func main() {
	router := routes.RegisterRoutes()

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
