package routers

import (
	"errors"
	"strings"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var UserID int
var Email string

func ProccessToken(tk string) (*models.Claims, bool, error) {
	myKey := []byte("Power_Metal_83")
	claims := &models.Claims{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, errors.New("split error")
	}

	token := strings.TrimSpace(splitToken[1])

	tkr, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := db.CheckEmail(claims.Email)
		if found {
			UserID = claims.ID
			Email = claims.Email
		}

		return claims, true, nil
	}

	if !tkr.Valid {
		return claims, false, errors.New("not valid token")
	}

	return claims, false, err
}
