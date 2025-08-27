package httpserver

import (
	"game/pkg/richerror"
	"game/servis"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (s Server) UserProfileHandler(c echo.Context) error {
	const op = "httpserver.UserProfileHandler"

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
		msg, code := CodeAndMessage(pErr)
		return echo.NewHTTPError(code, msg)
	}
 
	return c.JSON(http.StatusOK, response)
}

func CodeAndMessage(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		return re.Massage(), MapKindToHTTPStatusCode(re.Kind())
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func MapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindUnExepted:
		return http.StatusInternalServerError
	}

	return http.StatusBadRequest
}
