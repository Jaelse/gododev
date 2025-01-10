package repo

import (
	"fmt"

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
	Create(schedule *models.Schedule) models.Schedule
	GetNextSchedules() []models.Schedule
	MarkIsDone(id uint) models.Schedule
}

func (sc ScheduleRepo) Create(schedule *models.Schedule) models.Schedule {
	result := sc.db.Create(schedule)
	if result.Error != nil {
		fmt.Errorf("Error while creating schedule: %s", result.Error.Error())
	}
	return *schedule
}

func (sc ScheduleRepo) GetNextSchedules() []models.Schedule {
	var schs []models.Schedule
	sc.db.Find(&schs)
	return schs
}

func (sc ScheduleRepo) MarkIsDone(id uint) models.Schedule {
	var schedule models.Schedule
	sc.db.First(schedule, id)

	schedule.IsDone = true

	sc.db.Save(&schedule)
	return schedule
}
