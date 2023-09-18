package handlers

import (
	"encoding/json"
	externals "hp-api/externals/characters"
	"hp-api/models"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var character models.Character

	character, err := externals.GetCharacterById(id)
	if err != nil {
		log.Printf("Error to get the data: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}
