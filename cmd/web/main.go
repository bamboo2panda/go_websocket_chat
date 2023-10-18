package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

func main() {

	mux := routes()

	log.Println("Starting web server on port 8081")
	go handlers.ListenToWsChannel()

	_ = http.ListenAndServe(":8081", mux)
}
