package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"split-that.com/split-that/v2/src/logger"
)

func GetClaims(jwtToken string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})

	if err != nil {
		logger.Error.Println("Error parsing JWT Token", err.Error())
		return nil, errors.New("error parsing JWT Token")
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		logger.Error.Println("Invalid JWT token")
		return nil, errors.New("invalid JWT token")
	}
	return claims, nil
}
