package main

import (
	"html/template"
	"io"
	"log"

	"github.com/gododev/internal/presentation"
	"github.com/gododev/internal/repo"
	"github.com/gododev/pkg/do"
	"github.com/gododev/pkg/scheduler"
	"github.com/gododev/pkg/sqlitedb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	db := dbstuff()

	var scheduleRepo repo.IScheduleRepo
	scheduleRepo = repo.NewScheduleRepo(db)

	var snapshotRepo repo.ISnapshotRepo
	snapshotRepo = repo.NewSnapshotRepo(db)

	doc := do.NewClient()
	sch := scheduler.New(scheduleRepo,snapshotRepo)
	sch.Run()

	view(scheduleRepo, doc)
}

func dbstuff() *gorm.DB {
	db, err := sqlitedb.ConfigureDB()
	if err != nil {
		log.Fatal("error configuring Sqlite: \n %s", err.Error())
	}

	err = sqlitedb.Migrate(db)
	if err != nil {
		log.Fatal("Error while auto migrating: \n %s")
	}

	return db
}

func view(scheduleRepo repo.IScheduleRepo, doc do.DoClient) {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()
	presentation.LoadIndex(e)
	presentation.LoadSchedules(e, scheduleRepo)
	presentation.LoadDroplets(e, doc)
	e.Logger.Fatal(e.Start(":42069"))
}
