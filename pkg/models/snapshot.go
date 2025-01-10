package models

import "gorm.io/gorm"

type Snapshot struct {
	gorm.Model
	Name string
	Drop string
}
