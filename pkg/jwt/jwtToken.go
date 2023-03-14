package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var nowDate = time.Now().Format("2006-01-02 15")

func GenerateToken(mapClaims jwt.MapClaims, key string) (string, error) {
	key = fmt.Sprintf("%v%v", nowDate, key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, mapClaims)
	return token.SignedString([]byte(key))
}

func ParseToken(token, key string) (string, error) {
	key = fmt.Sprintf("%v%v", nowDate, key)
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	return claims.Claims.(jwt.MapClaims)["cmd"].(string), nil
}
