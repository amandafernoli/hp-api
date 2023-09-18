package handlers

import (
	"encoding/json"
	externals "hp-api/externals/spells"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	Spell, err := externals.ListSpells()

	if err != nil {
		log.Printf("Error to get the data: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Spell)
}
