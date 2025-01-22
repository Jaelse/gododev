package presentation

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gododev/internal/repo"
	"github.com/gododev/pkg/models"
	"github.com/labstack/echo/v4"
)

func LoadSchedules(e *echo.Echo, scheduleRepo repo.IScheduleRepo) {
	e.GET("/schedules", func(c echo.Context) error {
		schs := scheduleRepo.GetNextSchedules()
		return c.Render(200, "schedules", struct{ Schedules []models.DownSchedule }{Schedules: schs})
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

		att, err := convertToNextTime(at)
		if err != nil {
			return err
		}

		_ = scheduleRepo.Create(&models.DownSchedule{
			DropletID: uint(dropletID),
			At:        *att,
			IsDone:    false,
			Repeat:    repeat,
		})

		schs := scheduleRepo.GetNextSchedules()
		return c.Render(200, "schedules", struct{ Schedules []models.DownSchedule }{Schedules: schs})
	})

	e.DELETE("/schedules/:id", func(c echo.Context) error {

		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 42)
		if err != nil {
			return err
		}

		scheduleRepo.Delete(uint(id))

		schs := scheduleRepo.GetNextSchedules()
		return c.Render(200, "schedules", struct{ Schedules []models.DownSchedule }{Schedules: schs})
	})
}

func convertToNextTime(at string) (*time.Time, error) {
	parts := strings.Split(at, ":")

	// Validate that the split resulted in exactly two parts
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid input format: expected 'hh:mm'")
	}

	hr, err := strconv.ParseUint(parts[0], 10, 42)

	if err != nil {
		return nil, err
	}

	mm, err := strconv.ParseUint(parts[1], 10, 42)

	if err != nil {
		return nil, err
	}

	curr := time.Now()

	var newTime time.Time

	if curr.Hour() > int(hr) {
		newTime = time.Date(curr.Year(), curr.Month(), curr.Day()+1, int(hr), int(mm), 0, 0, curr.Location())
	}
	newTime = time.Date(curr.Year(), curr.Month(), curr.Day(), int(hr), int(mm), 0, 0, curr.Location())

	return &newTime, nil
}
