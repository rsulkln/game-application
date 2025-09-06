package main

import (
	"fmt"
	"game/auth"
	"game/config"
	_const "game/const"
	"game/delivery/httpserver"
	"game/repository/migrator"
	"game/repository/mysql"
	userservice "game/servis"
	"game/validator/uservalidator"
	"os"
	"strconv"
)

func getHttpServerPort(fallback int) int {
	/*The getHttpPortToInteger function retrieves the GAMEAPLICATION_PORT value from the environment.
	and converts it to an integer. If not found, it accepts the input value as the port.*/
	port := os.Getenv("GAMEAPPLICATION_PORT")

	outportt, err := strconv.Atoi(port)
	if err != nil {
		return fallback
	}
	fmt.Println(outportt)
	return outportt
}

func main() {
	cfg2 := config.Load("config.yml")
	fmt.Printf("cfg2: %+v\n ", cfg2)

	cfg := config.Config{
		HTTPServerConfig: config.HTTPServerConfig{getHttpServerPort(8012)},
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

	authsvc, usersvg, userValidator := setUpServices(cfg)
	server := httpserver.New(cfg, authsvc, usersvg, userValidator)
	server.Serve()
}

func setUpServices(cfg config.Config) (auth.Serivce, userservice.Service, uservalidator.Validator) {
	authSvc := auth.New(cfg.Auth)
	mysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, mysqlRepo)
	uv := uservalidator.New(mysqlRepo)

	return authSvc, userSvc, uv
}
