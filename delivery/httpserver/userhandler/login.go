package userhandler

import (
	"game/param"
	"game/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) LoginHandler(c echo.Context) error {
	var lReq param.LoginRequest
	if err := c.Bind(&lReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)

	}

	if fieldsError, err := h.userValidator.LoginValidationRequest(lReq); err != nil {
		msg, code := httpmsg.CodeAndMessage(err)

		return c.JSON(code, echo.Map{
			"message":     msg,
			"error field": fieldsError,
		})
	}

	resp, lErr := h.userSvc.Login(lReq)
	if lErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.JSON(http.StatusOK, resp); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)

	}
	return nil
}
