package presentation

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/gododev/pkg/do"
	"github.com/labstack/echo/v4"
)

func LoadDroplets(e *echo.Echo, doc do.DoClient) {
	e.GET("/droplets", func(c echo.Context) error {
		droplets, err := doc.GetAll(context.TODO())

		if err != nil {
			return c.Render(500, "Error", 0)
		}
		return c.Render(200, "droplets", struct{ Droplets []godo.Droplet }{Droplets: droplets})
	})
}
