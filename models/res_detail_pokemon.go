package models

type ResDetailPokemon struct {
	Abilities []ResAbility `json:"abilities"`
	Height    int          `json:"height"`
	Name      string       `json:"name"`
	Species   ResSpecies   `json:"species"`
	Weight    int          `json:"weight"`
	PokemonId int          `json:"id"`
}
