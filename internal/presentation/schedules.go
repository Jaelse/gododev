package presentation

import (
	"fmt"

	"github.com/gododev/internal/repo"
	"github.com/gododev/pkg/models"
	"github.com/labstack/echo/v4"
)

func LoadSchedules(e *echo.Echo, scheduleRepo repo.IScheduleRepo) {
	e.GET("/schedules", func(c echo.Context) error {
		schs := scheduleRepo.GetNextSchedules()
		fmt.Println(len(schs))
		return c.Render(200, "schedules", struct{ Schedules []models.Schedule }{Schedules: schs})
	})

	e.POST("/schedules", func(c echo.Context) error {
		_ = scheduleRepo.Create(&models.Schedule{
			DropletID: 112121212,
			IsDone:    false,
			Repeat:    true,
		})

		schs := scheduleRepo.GetNextSchedules()
		return c.Render(200, "schedules", struct{ Schedules []models.Schedule }{Schedules: schs})
	})
}
