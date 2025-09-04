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
	authConfig    auth.Config
}

func New(authConfig auth.Config, authSvc auth.Serivce, userSvc userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authConfig:    authConfig,
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
