package models

type WizardWand struct {
	Wood   string
	Core   string
	length int
}

type Character struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	AlternateNames   []string   `json:"alternate_names"`
	Species          string     `json:"species"`
	Gender           string     `json:"gender"`
	House            string     `json:"house"`
	DateOfBirth      string     `json:"dateOfBirth"`
	YearOfBirth      int        `json:"yearOfBirth"`
	Wizard           bool       `json:"wizard"`
	Ancestry         string     `json:"ancestry"`
	EyeColour        string     `json:"eyeColour"`
	HairColour       string     `json:"hairColour"`
	Wand             WizardWand `json:"wand"`
	Patronus         string     `json:"patronus"`
	HogwartsStudent  bool       `json:"hogwartsStudent"`
	HogwartsStaff    bool       `json:"hogwartsStaff"`
	Actor            string     `json:"actor"`
	Alternate_actors []string   `json:"alternate_actors"`
	Alive            bool       `json:"alive"`
	Image            string     `json:"image"`
}

type Spell struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
