package httpserver

import (
	"game/servis"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) LoginHandler(c echo.Context) error {
	var lReq servis.LoginRequest
	if err := c.Bind(&lReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)

	}
	resp, lErr := s.userSvc.Login(lReq)
	if lErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.JSON(http.StatusOK, resp); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)

	}
	return nil
}
