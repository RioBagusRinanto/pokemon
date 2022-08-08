package models

import (
	"time"
)

type DetailSpecieses struct {
	Species_id int
	Pokemon_id int
	CreateAt   time.Time
	UpdateAt   time.Time
}
