package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(mapClaims jwt.MapClaims, key string, Expire int64) (string, error) {
	iat := time.Now().Unix()
	mapClaims["exp"] = iat + Expire
	mapClaims["iat"] = iat
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, mapClaims)
	return token.SignedString([]byte(key))
}

func ParseToken(token, key string) (string, error) {
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	return claims.Claims.(jwt.MapClaims)["cmd"].(string), nil
}
