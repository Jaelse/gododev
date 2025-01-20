package scheduler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/digitalocean/godo"
	"github.com/gododev/internal/repo"
	"github.com/gododev/pkg/do"
	"github.com/gododev/pkg/models"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	dropSrv      do.DropletSrv
	scheduleRepo repo.IScheduleRepo
	SnapshotRepo repo.ISnapshotRepo
}

func New(scheduleRepo repo.IScheduleRepo, snapshotRepo repo.ISnapshotRepo) Scheduler {
	var droplet do.DropletSrv = do.NewClient()

	return Scheduler{
		dropSrv:      droplet,
		scheduleRepo: scheduleRepo,
		SnapshotRepo: snapshotRepo,
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
	schs := s.scheduleRepo.GetNextSchedules()

	if len(schs) == 0 {
		log.Println("No schedules")
		return
	}

	for _, sc := range schs {
		//TODO: check if the schedule has passed atleast 1 min or is stale for 12 hours
		if time.Now().Before(sc.At) {
			log.Printf("Still have time to take snapshot\n")
			break
		}

		dplt, err := s.dropSrv.Get(int(sc.DropletID), context.TODO())
		if err != nil {
			log.Fatal("%s", err.Error())
		}

		// Check if the snapshot is initiated
		snapshot, _ := s.SnapshotRepo.GetByScheduleID(sc.ID)
		if snapshot == nil {
			s.initializeSnapshot(sc, dplt)
		} else {
			s.initializeDropletKill(sc, snapshot, dplt)
		}
	}
}

func (s Scheduler) initializeSnapshot(sc models.Schedule, dplt *godo.Droplet) {
	snapshotName := fmt.Sprintf("%s-%d%d%d-%d%d", dplt.Name, time.Now().Day(), time.Now().Month(), time.Now().Year(), time.Now().Hour(), time.Now().Minute())
	log.Printf("Taking snapshot for %s\n", dplt.Name)
	err := s.dropSrv.TakeSnapshop(int(sc.DropletID), snapshotName, context.TODO())
	if err != nil {
		log.Fatal("%s", err.Error())
	}

	s.SnapshotRepo.Create(&models.Snapshot{
		Name:       snapshotName,
		Drop:       sc.DropletID,
		IsReady:    false,
		ScheduleId: sc.ID,
	})

	log.Printf("Snapshot with name %s created for droplet %s\n", snapshotName, dplt.Name)
}

func (s Scheduler) initializeDropletKill(sc models.Schedule, snaprec *models.Snapshot, dplt *godo.Droplet) {
	// If the snapshot is initiated then check if the snapshot is ready
	snaps, err := s.dropSrv.ListSnapshotByDropeltId(int(sc.DropletID), context.TODO())
	if err != nil {
		fmt.Errorf("Error listing the snapshots for droplet %s\n", dplt.Name)
	}

	for _, snap := range snaps {
		log.Printf("Snapshot: %s\n", snap.Name)

		if snap.Name == snaprec.Name {
			_, err := s.SnapshotRepo.MarkAsReady(snaprec.ID)
			if err != nil {
				log.Fatal(err.Error())
			}
			// If the snapshot is ready then kill the droplet and
			s.dropSrv.Kill(int(sc.DropletID), context.TODO())
			// If the schedule was in repeate then
			if sc.Repeat {
				newTime := time.Date(sc.At.Year(), sc.At.Month(), sc.At.Day()+1, sc.At.Hour(), sc.At.Minute(), 0, 0, sc.At.Location())
				// make a new schedule, add the droplet id and mark it as repeat
				newsch := s.scheduleRepo.Create(&models.Schedule{
					DropletID: sc.DropletID,
					Repeat:    sc.Repeat,
					IsDone:    false,
					At:        newTime,
				})

				log.Printf("Created next schedule for droplet %d at %s", newsch.DropletID, newsch.At)
			}

			s.scheduleRepo.MarkIsDone(sc.ID)
			log.Printf("Marked schedule for droplet %d as done", sc.DropletID)
		}
	}
}
