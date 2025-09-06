package config

import _const "game/const"

var DefaultConfig = map[string]interface{}{
	"auth.access_subject":  _const.AccessTokenSubject,
	"auth.refresh_subject": _const.RefreshTokenSubject,
}
