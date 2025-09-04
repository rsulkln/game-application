package middleware

import (
	"game/auth"
	"game/const/delconst"
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Auth(service auth.Serivce, config auth.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey:    delconst.AuthDeliverConstKey,
		SigningKey:    []byte(config.Signkey),
		SigningMethod: "HS256",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claims, nil
		},
	})
}
