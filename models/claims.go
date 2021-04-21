package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID    int    `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}
