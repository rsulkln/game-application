package httpmsg

import (
	"game/const/errormessage"
	"game/pkg/richerror"
	"net/http"
)

func CodeAndMessage(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.Massage()
		code := mapKindToHTTPStatusCode(re.Kind())
		if code > 500 {
			msg = errormessage.SomthingWentWrong
		}

		return msg, code

	default:
		return err.Error(), http.StatusBadRequest
	}
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
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
