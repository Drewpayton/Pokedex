package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Location struct {
	EncounterMethodRates []struct{
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"` 
		} `json:"encounter_method"`
		VersionDetails struct {
			Rate    int `json:"rate"`
			Version struct{
				Name string `json:"name"`
				URL  string `json:"url"`
			}`json:"version"`
		} `json:"version_deatils"`
	}`json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID int `json:"id"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name string `json:"name"`
		Language struct{
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	}`json:"names"`
	PokemonEncounter []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}`json:"version"`
			MaxChance int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				ConditionValues []struct{} `json:"condition_values"`
				Chance int `json:"chance"`
				Method struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			}`json:"encounter_details"`
		}`json:"version_details"`
	} `json:"pokemon_encounters"`
}
