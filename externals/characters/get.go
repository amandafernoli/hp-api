package externals

import (
	"encoding/json"
	"fmt"
	"hp-api/configs"
	"hp-api/models"
	"io"
	"net/http"
)

func GetCharacterById(id string) (models.Character, error) {
	resp, err := http.Get(configs.GetBaseURL() + "/character/" + id)

	if err != nil {
		fmt.Println(err.Error())
		return models.Character{}, err
	}

	body, err := io.ReadAll(resp.Body)

	var character []models.Character
	err = json.Unmarshal(body, &character)

	if err != nil {
		fmt.Println("Error to get the data", err.Error())
		return models.Character{}, err
	}

	return character[0], nil
}

func GetCharactersByHouse(house string) ([]models.Character, error) {
	resp, err := http.Get(configs.GetBaseURL() + "/characters/house/" + house)

	if err != nil {
		fmt.Println(err.Error())
		return []models.Character{}, err
	}

	body, err := io.ReadAll(resp.Body)

	var characters []models.Character
	err = json.Unmarshal(body, &characters)

	if err != nil {
		fmt.Println("Error to get the data", err.Error())
		return []models.Character{}, err
	}

	return characters, nil
}

func ListCharacters() ([]models.Character, error) {
	resp, err := http.Get(configs.GetBaseURL() + "/characters")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	var characters []models.Character
	err = json.Unmarshal(body, &characters)

	if err != nil {
		fmt.Println("Error to get the data", err.Error())
		return nil, err
	}

	return characters, nil
}
