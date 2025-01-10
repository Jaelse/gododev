package presentation

import "github.com/labstack/echo/v4"

func LoadIndex(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", 0)
	})
}
