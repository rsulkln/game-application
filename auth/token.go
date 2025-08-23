package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func (s Serivce) CreateToken(userid uint, subject string, ExpireDuration time.Duration) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpireDuration)),
		},
		UserID: userid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString([]byte(s.config.Signkey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
