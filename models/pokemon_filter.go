package models

type PokemonFilter struct {
	Name    []string `json:"name"`
	Weight  string   `json:"weight"`
	Height  string   `json:"height"`
	Ability []string `json:"ability"`
	Species []string `json:"species"`
}
