package userhandler

import (
	"game/dto"
	"game/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) UserProfileHandler(c echo.Context) error {
	const op = "httpserver.UserProfileHandler"

	authToken := c.Request().Header.Get("Authorization")

	if authToken == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is empty")
	}

	claim, pErr := h.authSvc.ParseToken(authToken)
	if pErr != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, pErr.Error())
	}

	response, pErr := h.userSvc.Profile(dto.ProfileRequest{UserID: claim.UserID})
	if pErr != nil {
		msg, code := httpmsg.CodeAndMessage(pErr)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, response)
}
