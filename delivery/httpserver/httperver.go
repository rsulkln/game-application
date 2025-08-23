package httpserver

import (
	"fmt"
	"game/auth"
	"game/config"

	userservice "game/servis"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config  config.Config
	authSvc auth.Serivce
	userSvc userservice.Service
}

func New(config config.Config, authSvc auth.Serivce, userSvc userservice.Service) Server {
	return Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}

func (s Server) Serve() {
	//start engine
	e := echo.New()

	//middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routers

	userGroup := e.Group("/users")
	userGroup.POST("/register", s.UserRegisterHandler)
	userGroup.POST("/login", s.LoginHandler)
	userGroup.GET("/profile", s.UserProfileHandler)

	e.GET("/health-check", s.HealthCheck)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServerConfig.Port)))
}
