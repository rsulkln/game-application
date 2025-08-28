package uservalidator

import (
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

func (v Validator) RegisteValidationRequest(req dto.RegisterRequest) error {
	const op = "uservalidator.RegisteValidationRequest"
	var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#\$%\^&\*])[A-Za-z\d!@#\$%\^&\*]{8,}$`)

	if vErr := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(passwordRegex)),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(`^09[0-9]{9}$`))),
	); vErr != nil {

		return richerror.
			New(op).
			WithMassage(errormessage.InvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithError(vErr)

	}

	if isUniq, err := v.repo.IsUniquePhoneNumber(req.PhoneNumber); err != nil || !isUniq {
		if err != nil {
			richerror.New(op).WithError(err)
		}
		if !isUniq {
			return richerror.
				New(op).
				WithMassage(errormessage.PhoneNumberNotUniq).
				WithKind(richerror.KindInvalid).
				WithMeta(map[string]interface{}{"phone Number": req.PhoneNumber})
		}
	}
}
