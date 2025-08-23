package httpserver

import (
	"game/servis"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) UserProfileHandler(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	if authToken == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is empty")
	}
	claim, pErr := s.authSvc.ParseToken(authToken)
	if pErr != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, pErr.Error())
	}
	response, pErr := s.userSvc.Profile(servis.ProfileRequest{UserID: claim.UserID})
	if pErr != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, pErr.Error())
	}

	return c.JSON(http.StatusOK, response)
}
