package main

import (
	"fmt"
	"hp-api/configs"
	"hp-api/externals"
	characters "hp-api/handlers/characters"
	spells "hp-api/handlers/spells"
	"hp-api/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	utils.PrintWelcome()

	err := externals.ConnectApi()
	if err != nil {
		fmt.Println("Error connecting to API:", err)
	}

	r := chi.NewRouter()
	r.Get("/{id}", characters.Get)
	r.Get("/characters", characters.List)
	r.Get("/characters/house/{house}", characters.ListByHouse)
	r.Get("/spells", spells.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
