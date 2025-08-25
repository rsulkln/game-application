package main

import (
	"game/auth"
	"game/config"
	_const "game/const"
	"game/delivery/httpserver"
	"game/repository/migrator"

	"game/repository/mysql"
	userservice "game/servis"
)

func main() {
	cfg := config.Config{
		HTTPServerConfig: config.HTTPServerConfig{Port: 8088},
		Auth: auth.Config{
			Signkey:           _const.JwtSignKey,
			AccessExpireTime:  _const.AccessExpireTime,
			RefreshExpireTime: _const.RefreshExpireTime,
			AccessSubject:     _const.AccessTokenSubject,
			RefreshSubject:    _const.RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: "rasool",
			Password: "60BA944D1AACA915C803676D11C105A2",
			Host:     "localhost",
			Port:     3306,
			Database: "game_application",
		},
	}

	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	authsvc, usersvg := setUpServices(cfg)
	server := httpserver.New(cfg, authsvc, usersvg)
	server.Serve()
}

func setUpServices(cfg config.Config) (auth.Serivce, userservice.Service) {
	authSvc := auth.New(cfg.Auth)
	mysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, mysqlRepo)
	return authSvc, userSvc
}
