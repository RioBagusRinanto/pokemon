package models

import (
	"time"
)

type Abilities struct {
	Id           int
	Ability_Name string
	Url          string
	CreateAt     time.Time
	UpdateAt     time.Time
}
