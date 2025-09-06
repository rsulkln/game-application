package config

import (
	"game/auth"
	"game/repository/mysql"
)

type HTTPServerConfig struct {
	Port int `koanf:"port"`
}
type Config struct {
	HTTPServerConfig HTTPServerConfig `koanf:"http_server"`
	Auth             auth.Config      `koanf:"auth"`
	Mysql            mysql.Config     `koanf:"mysql"`
}
