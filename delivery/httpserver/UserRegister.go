package httpserver

import (
	"game/dto"
	"game/pkg/httpmsg"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (s Server) UserRegisterHandler(c echo.Context) error {
	var uReq dto.RegisterRequest

	if bErr := c.Bind(&uReq); bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if fieldsError, err := s.userValidator.RegisteValidationRequest(uReq); err != nil {
		msg, code := httpmsg.CodeAndMessage(err)

		return c.JSON(code, echo.Map{
			"message":     msg,
			"error field": fieldsError,
		})
	}

	response, rErr := s.userSvc.Register(uReq)
	if rErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, response)
}
