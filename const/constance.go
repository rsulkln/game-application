package _const

import "time"

const (
	JwtSignKey          = "9ACABEAC10A803749D13E10269505776"
	AccessTokenSubject  = "at"
	RefreshTokenSubject = "rt"
	AccessExpireTime    = time.Duration(24 * time.Hour)
	RefreshExpireTime   = time.Duration(24 * 7 * time.Hour)
)
