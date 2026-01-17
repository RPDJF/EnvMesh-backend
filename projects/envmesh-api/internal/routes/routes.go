package routes

import (
	"fmt"
	"net/http"

	"github.com/RPDJF/EnvMesh-backend/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	fmt.Println("Register Routes")
	router := mux.NewRouter()
	router.HandleFunc("/ping", handlers.PingHandler).Methods(http.MethodGet)
	return router
}
