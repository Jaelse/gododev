package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	DropletID uint
	Repeat    bool
	IsDone    bool
	At      time.Time
}
