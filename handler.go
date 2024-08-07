package main

import (
	"log"
	"net/http"
)

func handleGetNotes(w http.ResponseWriter, _ *http.Request) {
	notes, err := fetchNotes()
	if err != nil {
		log.Printf("Error fetching notes %v", err)
		http.Error(w, "Unable to fetch notes", http.StatusInternalServerError)
		return
	}
	log.Printf("Notes: %v", notes)
	err = tmpl.ExecuteTemplate(w, "Base", map[string]interface{}{"notes": notes})
	if err != nil {
		log.Printf("Error executing template %v", err)
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
