package utils

import (
	"hp-api/models"

	color "github.com/fatih/color"
)

const (
	GRYFFINDOR = "Gryffindor"
	SLYTHERIN  = "Slytherin"
	HUFFLEPUFF = "Hufflepuff"
	RAVENCLAW  = "Ravenclaw"
)

func PrintWelcome() {
	welcome := color.New(color.Bold)

	color.New(color.FgYellow).Print("Welcome, students of ")
	welcome.Add(color.FgRed).Print("Ho")
	welcome.Add(color.FgGreen).Print("gw")
	welcome.Add(color.FgYellow).Print("ar")
	welcome.Add(color.FgBlue).Print("ts")
	color.New(color.FgYellow).Println("!")
}

func PrintCharactersByHouse(characters []models.Character) {
	color.White("\nListing Harry Potter characters:")

	for _, character := range characters {
		if character.House == GRYFFINDOR {
			color.Red(character.Name)
		} else if character.House == HUFFLEPUFF {
			color.Yellow(character.Name)
		} else if character.House == RAVENCLAW {
			color.Blue(character.Name)
		} else if character.House == SLYTHERIN {
			color.Green(character.Name)
		} else if character.House == "" {
			color.White(character.Name)
		}
	}
}
