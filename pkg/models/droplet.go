package models

import "gorm.io/gorm"

type Droplet struct {
	gorm.Model
	DropletID uint
	Name      string
	Ip        string
	Snapshot  Snapshot
}
