package models

import (
	"time"
)

type Specieses struct {
	Id           int
	Species_Name string
	Url          string
	CreateAt     time.Time
	UpdateAt     time.Time
}
