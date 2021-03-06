package main

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type UserInfo struct {
	Id       uint   `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	TokenType string
	UserInfo
}

func ParseJWT(tokenString string, verifyKey *rsa.PublicKey) (UserInfo, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		return UserInfo{}, nil
	}

	if !token.Valid {
		return UserInfo{}, fmt.Errorf("invalid jwt")
	}
	claims := token.Claims.(*CustomClaims)

	return claims.UserInfo, nil
}

func generateJWT(info UserInfo, signKey *rsa.PrivateKey) (string, error) {

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims = &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
		"level1",
		info,
	}

	return token.SignedString(signKey)
}
