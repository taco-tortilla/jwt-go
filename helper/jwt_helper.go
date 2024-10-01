package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
}

func IsTokenExpired(claims jwt.MapClaims) bool {
	return float64(time.Now().Unix()) > claims["exp"].(float64)
}
