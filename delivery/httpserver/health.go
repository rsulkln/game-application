package httpserver

import (
	"github.com/labstack/echo/v4"
)

func (s Server) HealthCheck(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"massage": "every thing is good",
	})
}
