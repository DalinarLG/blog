package db

import "golang.org/x/crypto/bcrypt"

func Encrypt(pass string) string {

	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(bytes)
}
