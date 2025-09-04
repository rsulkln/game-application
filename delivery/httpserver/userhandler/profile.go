package userhandler

import (
	"fmt"
	"game/auth"
	"game/const/delconst"
	"game/param"
	"game/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func parseClaims(c echo.Context) *auth.Claims {
	const op = "httpserver.UserProfileHandler"

	claims := c.Get(delconst.AuthDeliverConstKey)
	cl, ok := claims.(*auth.Claims)
	if ok {
		fmt.Println("Successfully")
	} else {
		panic("claim in not found!")
	}
	return cl
}

func (h Handler) UserProfileHandler(c echo.Context) error {
	claims := parseClaims(c)

	response, pErr := h.userSvc.Profile(param.ProfileRequest{UserID: claims.UserID})
	if pErr != nil {
		msg, code := httpmsg.CodeAndMessage(pErr)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, response)
}
