package repo

import (
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
}

func (sr SnapshotRepo) Create(snap *models.Snapshot) models.Snapshot {
	sr.db.Create(snap)
	return *snap
}

func (sr SnapshotRepo) Get(id uint) models.Snapshot {
	var snapshot models.Snapshot
	sr.db.First(snapshot, id)
	return snapshot
}
