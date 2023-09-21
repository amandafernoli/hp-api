package handlers

import (
	"encoding/json"
	externals "hp-api/externals/characters"
	"hp-api/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func List(w http.ResponseWriter, r *http.Request) {
	characters, err := externals.ListCharacters()

	if err != nil {
		log.Printf("Error to get the data: %v", err)
	}

	utils.PrintCharactersByHouse(characters)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func ListByHouse(w http.ResponseWriter, r *http.Request) {
	house := chi.URLParam(r, "house")

	characters, err := externals.GetCharactersByHouse(house)

	if err != nil {
		log.Printf("Error to get the data: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}
