package scheduler

import (
	"context"
	"fmt"
	"log"

	"github.com/gododev/internal/repo"
	"github.com/gododev/pkg/do"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	dropSrv      do.DropletSrv
	scheduleRepo repo.IScheduleRepo
}

func New(scheduleRepo repo.IScheduleRepo) Scheduler {
	var droplet do.DropletSrv
	droplet = do.NewClient()

	return Scheduler{
		dropSrv:      droplet,
		scheduleRepo: scheduleRepo,
	}
}

func (s Scheduler) Run() {
	c := cron.New()

	c.AddFunc("@every 00h00m10s", s.dropDropletRoutine)

	c.Start()
	// run until stopped
	// run drop droplet routine according to schedule
	// run spinUpDroplet routine when schedule
}

func (s Scheduler) Stop() error {
	// run until stopped
	// run drop droplet routine according to schedule
	// run spinUpDroplet routine when schedule
	return nil
}

func (s Scheduler) dropDropletRoutine() {
	fmt.Println("Checking schedule for droplets")
	schs := s.scheduleRepo.GetNextSchedules()

	if len(schs) == 0 {
		fmt.Println("No schedules")
		return
	}
	for _, sc := range schs {
		dplt, err := s.dropSrv.Get(int(sc.DropletID), context.TODO())
		if err != nil {
			log.Fatal("%s", err.Error())
		}

		// err = s.dropSrv.TakeSnapshop(int(sc.DropletID), context.TODO())
		// if err != nil {
		// 	log.Fatal("%s", err.Error())
		// }
		fmt.Printf("Snapshot for droplet %d - %s is in progress", dplt.ID, dplt.Name)
	}
	// record the snapshot history
	// once snapshot is ready, kill the droplet
	//
}

func snipDropletRoutine() {
	// record the snapshot history
	// once snapshot is ready, kill the droplet
	//
}
