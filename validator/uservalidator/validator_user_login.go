package uservalidator

import (
	"fmt"
	"game/const/errormessage"
	"game/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"

	"game/pkg/richerror"
)

func (v Validator) LoginValidationRequest(req dto.LoginRequest) (map[string]string, error) {
	const op = "uservalidator.RegisteValidationRequest"

	if vErr := validation.ValidateStruct(&req,
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(PhoneNUmberRegex))),
	); vErr != nil {

		FieldError := make(map[string]string)
		errV, ok := vErr.(validation.Errors)
		if ok {
			for key, val := range errV {
				if val != nil {
					FieldError[key] = val.Error()
				}
			}
		}

		return FieldError, richerror.
			New(op).
			WithMassage(errormessage.InvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithError(vErr)
	}

	return nil, nil
}

func (v Validator) DosNumberIsExists(value interface{}) error {
	phoneNumber := value.(string)
	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		fmt.Errorf(errormessage.NotFound)
	}
	return nil
}
