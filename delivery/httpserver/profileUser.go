package httpserver

import (
	"game/servis"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) UserProfileHandler(c echo.Context) error {
	var uReq servis.ProfileRequest
	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	profileResponse, pErr := s.userSvc.Profile(uReq)
	if pErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, profileResponse)

}
