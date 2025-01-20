package models

import "gorm.io/gorm"

type Snapshot struct {
	gorm.Model
	Name string
	Drop uint
	IsReady bool
	ScheduleId uint
}
