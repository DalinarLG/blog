package jwt

import (
	"time"

	"github.com/DalinarLG/blog/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(u models.User)(string, error){
	myKey := []byte("Power_Metal_83")
	payload := jwt.MapClaims{

		"id": u.ID,
		"name":u.Name,
		"email":u.Email,
		"avatar":u.Avatar,
		"exp": time.Now().Add(24*time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(myKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}