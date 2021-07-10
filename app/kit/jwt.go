package kit

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtDate struct {
	AccountID int32
	CompanyID int32
	EmployID  int32
}

// 自定义Claims
type JwtClaims struct {
	Data JwtDate
	jwt.StandardClaims
}

const MAX_AGE = 24 * 30 // 保持1个月
var loginTimeOut = time.Duration(MAX_AGE) * time.Hour

func JwtCreate(data JwtDate) (string, error) {
	claims := &JwtClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(loginTimeOut).Unix(),
			Issuer:    Env.JwtIss,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Env.Jwt)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func JwtParse(tokenString string) (JwtDate, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return Env.Jwt, nil
	})
	if err != nil {
		return JwtDate{}, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims.Data, nil
	} else {
		return JwtDate{}, fmt.Errorf("token parse to JwtClaims error")
	}
}
