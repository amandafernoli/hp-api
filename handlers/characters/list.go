package handlers

import (
	"encoding/json"
	externals "hp-api/externals/characters"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	Character, err := externals.ListCharacters()

	if err != nil {
		log.Printf("Error to get the data: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Character)
}
