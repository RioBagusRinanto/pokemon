package models

// type ListPokemons struct {
// 	Id           int    `json:"id"`
// 	Pokemon_Name string `json:"pokemon_name"`
// 	Height       int    `json:"height"`
// 	Weight       int    `json:"weight"`
// 	Abilities    []struct {
// 		Ability_Name string `json:"ability_name"`
// 	} `json:"ability_name"`
// 	Species struct {
// 		Species_Name string `json:"species_name"`
// 	} `json:"species_name"`
// }

type ListPokemons struct {
	Id           int    `json:"id"`
	Pokemon_Name string `json:"pokemon_name"`
	Height       int    `json:"height"`
	Weight       int    `json:"weight"`
	Ability_Name string `json:"ability_name"`
	Species_Name string `json:"species_name"`
}
