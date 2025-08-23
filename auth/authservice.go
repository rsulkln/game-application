package auth

import (
	"game/entity"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type Config struct {
	Signkey           string
	AccessExpireTime  time.Duration
	RefreshExpireTime time.Duration
	AccessSubject     string
	RefreshSubject    string
}

type Serivce struct {
	config Config
}

func New(cfg Config) Serivce {
	return Serivce{
		config: cfg,
	}
}

func (s Serivce) CreateAccessToken(user entity.User, subject string) (string, error) {
	return s.CreateToken(user.ID, s.config.AccessSubject, s.config.AccessExpireTime)
}

func (s Serivce) CreateRefreshToken(user entity.User, subject string) (string, error) {
	return s.CreateToken(user.ID, s.config.RefreshSubject, s.config.RefreshExpireTime)
}

func (s Serivce) ParseToken(bearerToken string) (*Claims, error) {
	bearerToken = strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(bearerToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.Signkey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
