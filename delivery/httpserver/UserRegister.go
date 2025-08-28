package httpserver

import (
	userservice "game/servis"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (s Server) UserRegisterHandler(c echo.Context) error {
	var uReq userservice.RegisterRequest

	if bErr := c.Bind(&uReq); bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	response, rErr := s.userSvc.Register(uReq)
	if rErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, response)
}
