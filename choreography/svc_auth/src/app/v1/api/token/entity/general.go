package entity

import "github.com/dgrijalva/jwt-go"

type AuthDataModels struct {
	Level string `json:"level"`
}

type Claims struct {
	Session string `json:"session"`
	jwt.StandardClaims
}
