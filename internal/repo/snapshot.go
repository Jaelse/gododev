package repo

import (
	"fmt"

	"github.com/gododev/pkg/models"
	"gorm.io/gorm"
)

type SnapshotRepo struct {
	db *gorm.DB
}

func NewSnapshotRepo(db *gorm.DB) *SnapshotRepo {
	return &SnapshotRepo{
		db: db,
	}
}

type ISnapshotRepo interface {
	Create(snap *models.Snapshot) models.Snapshot
	Get(id uint) models.Snapshot
	GetByScheduleID(id uint) (*models.Snapshot, error)
	MarkAsReady(id uint) (*models.Snapshot, error)
}

func (sr SnapshotRepo) Create(snap *models.Snapshot) models.Snapshot {
	result := sr.db.Create(snap)
	if result.Error != nil {
		fmt.Errorf("Error while creating snapshot: %s", result.Error.Error())
	}
	return *snap
}

func (sr SnapshotRepo) Get(id uint) models.Snapshot {
	var snapshot models.Snapshot
	sr.db.First(snapshot, id)

	return snapshot
}

func (sr SnapshotRepo) GetByScheduleID(id uint) (*models.Snapshot, error) {
	var snap models.Snapshot
	res := sr.db.Where(&models.Snapshot{ScheduleId: id}).First(&snap)

	if res.Error != nil {
		return nil, res.Error
	}
	return &snap, nil
}

func (sr SnapshotRepo) MarkAsReady(id uint) (*models.Snapshot, error) {
	var snap models.Snapshot
	res := sr.db.First(&snap, id)
	if res.Error != nil {
		return nil, res.Error
	}

	snap.IsReady = true
	res = sr.db.Save(&snap)
	if res.Error != nil {
		return nil, res.Error
	}

	return &snap, nil
}
