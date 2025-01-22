package sqlitedb

import (
	"github.com/gododev/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SDB struct {
	Db *gorm.DB
}

func ConfigureDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("data/gododev.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Snapshot{}, &models.DownSchedule{})
	return err
}
