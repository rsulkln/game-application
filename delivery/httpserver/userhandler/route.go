package userhandler

import (
	"game/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoute(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.UserRegisterHandler)

	userGroup.POST("/login", h.LoginHandler)

	userGroup.GET("/profile", h.UserProfileHandler, middleware.Auth(h.authSvc, h.authConfig))
}
