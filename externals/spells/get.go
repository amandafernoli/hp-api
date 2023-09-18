package externals

import (
	"encoding/json"
	"fmt"
	"hp-api/configs"
	"hp-api/models"
	"io"
	"net/http"
)

func ListSpells() ([]models.Spell, error) {
	resp, err := http.Get(configs.GetBaseURL() + "/spells")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	var spells []models.Spell
	err = json.Unmarshal(body, &spells)

	if err != nil {
		fmt.Println("Error to get the data", err.Error())
		return nil, err
	}

	return spells, nil
}
