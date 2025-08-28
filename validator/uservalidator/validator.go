package uservalidator

import (
	"fmt"
	"game/const/errormessage"
	"regexp"

	"game/dto"
	"github.com/go-ozzo/ozzo-validation/v4"

	"game/pkg/richerror"
)

type Repository interface {
	IsUniquePhoneNumber(phoneNumber string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repository Repository) Validator {
	return Validator{repo: repository}

}

func (v Validator) RegisteValidationRequest(req dto.RegisterRequest) (error, map[string]string) {
	const op = "uservalidator.RegisteValidationRequest"
	var passwordRegex = regexp.MustCompile(`^[A-Za-z\d!@#\$%\^&\*]{8,}$`)

	if vErr := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(passwordRegex)),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(`^09[0-9]{9}$`)),
			validation.By(v.CheckPhonneNumberIsUniqeness)),
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

		return richerror.
			New(op).
			WithMassage(errormessage.InvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithError(vErr), FieldError
	}

	return nil, nil
}

func (v Validator) CheckPhonneNumberIsUniqeness(value interface{}) error {
	phoneNumber := value.(string)
	if isUniq, err := v.repo.IsUniquePhoneNumber(phoneNumber); err != nil || !isUniq {
		if err != nil {
			return err
		}
		if !isUniq {
			fmt.Errorf(errormessage.PhoneNumberNotUniq)
		}
	}
	return nil
}
