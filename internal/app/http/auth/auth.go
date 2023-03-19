package auth

import (
	"fengjx/go-web-quickstart/pkg/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	secret = []byte("my_web_app_hello")
)

type LoginClaims struct {
	Uid int64
	jwt.RegisteredClaims
}

func Signed(uid int64) (string, error) {
	now := utils.Now
	exp := now.Add(time.Hour * 24 * 7)
	claims := LoginClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Parse(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		return claims.Uid, nil
	} else {
		return 0, err
	}
}
