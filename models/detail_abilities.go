package models

import (
	"time"
)

type DetailAbilities struct {
	Abilities_id int
	Pokemon_id   int
	CreateAt     time.Time
	UpdateAt     time.Time
}
