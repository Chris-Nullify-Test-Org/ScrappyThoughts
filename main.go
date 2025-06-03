package main

import (
	"log"
	"net/http"

	"scrappythoughts.com/scrappythoughts-repo/internal/server"
)

func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	log.Printf("Server starting on :8085")
	if err := http.ListenAndServe(":8085", srv.Router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
