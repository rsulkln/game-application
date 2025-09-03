package userhandler

import (
	"game/auth"
	userservice "game/servis"
	"game/validator/uservalidator"
)

type Handler struct {
	authSvc       auth.Serivce
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authSvc auth.Serivce, userSvc userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
