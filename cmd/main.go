package main

import (
	"log"
	"net/http"

	"github.com/the-fourth-dimension/hyperspace/pkg/hub"
)

func main() {
	logger := log.Default()
	logger.Println("Server Started")
	hub := hub.NewHub()
	http.HandleFunc("/ws", hub.ServeWS)
	http.ListenAndServe(":8000", nil)
}
