package repo

import (
	"fmt"
	"log"

	"github.com/gododev/pkg/models"
	"gorm.io/gorm"
)

type ScheduleRepo struct {
	db *gorm.DB
}

func NewScheduleRepo(db *gorm.DB) *ScheduleRepo {
	return &ScheduleRepo{
		db: db,
	}
}

type IScheduleRepo interface {
	Create(schedule *models.DownSchedule) models.DownSchedule
	GetNextSchedules() []models.DownSchedule
	GetNextScheduleByDropletId(dropletId uint) models.DownSchedule
	MarkIsDone(id uint) models.DownSchedule
	UpdateSnapshot(id uint, snapshot models.Snapshot) models.DownSchedule
	Delete(ID uint) bool
}

func (sc ScheduleRepo) Create(schedule *models.DownSchedule) models.DownSchedule {
	result := sc.db.Create(schedule)
	if result.Error != nil {
		fmt.Errorf("Error while creating schedule: %s", result.Error.Error())
	}
	return *schedule
}

func (sc ScheduleRepo) GetNextSchedules() []models.DownSchedule {
	var schs []models.DownSchedule
	sc.db.Where("is_done = ?", false).Find(&schs)
	return schs
}

func (sc ScheduleRepo) GetNextScheduleByDropletId(dropletId uint) models.DownSchedule {
	var sch models.DownSchedule
	sc.db.Where(&models.DownSchedule{DropletID: dropletId, IsDone: false}).First(&sch)

	return sch
}

func (sc ScheduleRepo) UpdateSnapshot(id uint, snapshot models.Snapshot) models.DownSchedule {
	var sch models.DownSchedule
	sc.db.First(&sch, id)

	sch.Snapshot = snapshot

	sc.db.Save(&sch)

	return sch
}

func (sc ScheduleRepo) MarkIsDone(id uint) models.DownSchedule {
	var schedule models.DownSchedule

	res := sc.db.First(&schedule, id)

	if res.Error != nil {
		log.Fatal(res.Error.Error())
	}
	schedule.IsDone = true

	sc.db.Save(&schedule)
	return schedule
}

func (sc ScheduleRepo) Delete(ID uint) bool {
	result := sc.db.Delete(&models.DownSchedule{}, ID)

	if result.Error != nil {
		fmt.Errorf("Error while deleting schedule: %s", result.Error.Error())
		return false
	}

	return true
}
