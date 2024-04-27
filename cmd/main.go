package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/the-fourth-dimension/hyperspace/pkg/env"
	"github.com/the-fourth-dimension/hyperspace/pkg/hub"
)

func main() {
	env.LoadEnv()
	logger := log.Default()
	hub := hub.NewHub()
	http.HandleFunc("/ws", hub.ServeWS)
	logger.Printf("starting server on %s\n", env.GetEnv(env.PORT))
	err := http.ListenAndServe(fmt.Sprintf(":%s", env.GetEnv(env.PORT)), nil)
	if err != nil {
		logger.Fatalf("an error occurred while starting the server: %v", err)
		panic(1)
	}
}
