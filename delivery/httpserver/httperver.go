package httpserver

import (
	"fmt"
	"game/auth"
	"game/config"
	"game/delivery/httpserver/userhandler"
	userservice "game/servis"
	"game/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func New(config config.Config,
	authSvc auth.Serivce,
	userSvc userservice.Service,
	userValidator uservalidator.Validator,
) Server {
	return Server{
		config:      config,
		userHandler: userhandler.New(config.Auth, authSvc, userSvc, userValidator),
	}
}

func (s Server) Serve() {
	//start engine
	e := echo.New()

	//middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routers
	e.GET("/health-check", s.HealthCheck)
	s.userHandler.SetUserRoute(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServerConfig.Port)))
}
