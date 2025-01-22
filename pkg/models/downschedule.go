package models

import (
	"time"

	"gorm.io/gorm"
)

type DownSchedule struct {
	gorm.Model
	DropletID uint
	Repeat    bool
	IsDone    bool
	At        time.Time
	Snapshot  Snapshot
}
