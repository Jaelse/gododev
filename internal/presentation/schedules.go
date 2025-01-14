package presentation

import (
	"fmt"
	"strconv"

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
		dropletID, err := strconv.ParseUint(c.FormValue("schedule-droplet"), 10, 42)
		if err != nil {
			return err
		}

		at := c.FormValue("schedule-time")

		repeat, err := strconv.ParseBool(c.FormValue("repeat"))
		if err != nil {
			return err
		}

		_ = scheduleRepo.Create(&models.Schedule{
			DropletID: uint(dropletID),
			At:        at,
			IsDone:    false,
			Repeat:    repeat,
		})

		schs := scheduleRepo.GetNextSchedules()
		return c.Render(200, "schedules", struct{ Schedules []models.Schedule }{Schedules: schs})
	})

	e.DELETE("/schedules/:id", func(c echo.Context) error {

		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 42)
		if err != nil {
			return err
		}

		scheduleRepo.Delete(uint(id))

		schs := scheduleRepo.GetNextSchedules()
		return c.Render(200, "schedules", struct{ Schedules []models.Schedule }{Schedules: schs})
	})
}
