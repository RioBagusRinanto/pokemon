package models

import (
	"time"
)

type Pokemons struct {
	Id           int       `json:"id"`
	Pokemon_Name string    `json:"pokemon_name"`
	Height       int       `json:"height"`
	Weight       int       `json:"weight"`
	CreateAt     time.Time `json:"createAt"`
	UpdateAt     time.Time `json:"updateAt"`
}
