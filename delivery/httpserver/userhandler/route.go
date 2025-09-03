package userhandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoute(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.UserRegisterHandler)
	userGroup.POST("/login", h.LoginHandler)
	userGroup.GET("/profile", h.UserProfileHandler)
}
