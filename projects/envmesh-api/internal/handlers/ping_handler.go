package handlers

import (
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PingHandler call")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
