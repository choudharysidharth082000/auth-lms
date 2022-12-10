package mailer

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// function to mail the user
func MailUser(email string, subject string, body string) bool {
	//setting up the mailer
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//setting up the dialer
	d := gomail.NewDialer("sidharth@prodigalai.com", 587, "", "Sidharth@2000")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
