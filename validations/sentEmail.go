package validations

import (
	"log"
	"net/smtp"
)

func SendEmail(email string) error {
	from := "strongerlg@gmail.com"
	password := "ldcigglwrmoexcbs"
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//subject := []byte("Welcome")
	message := []byte("Welcome to my api")	

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
