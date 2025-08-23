package config

import (
	"game/auth"
	"game/repository/mysql"
)

type HTTPServerConfig struct {
	Port int
}
type Config struct {
	HTTPServerConfig HTTPServerConfig
	Auth             auth.Config
	Mysql            mysql.Config
}
