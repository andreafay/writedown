package main

import (
	"log"
	"net/http"
)

func main() {
	// Database setup
	err := openDB()
	if err != nil {
		log.Panic(err)
	}
	defer closeDB()
	err = setupDB()
	if err != nil {
		log.Panic(err)
	}
	err = parseTemplates()
	if err != nil {
		log.Panic(err)
	}

	// Server setup
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleGetNotes)

	log.Println("Server starting on localhost:3001")
	if err := http.ListenAndServe("localhost:3001", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
