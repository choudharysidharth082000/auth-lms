package mailer

import "gopkg.in/gomail.v2"

//function to mail the user
func MailUser(email string, subject string, body string) bool {
	//setting up the mailer
	m := gomail.NewMessage()
	m.SetHeader("From", email);
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//setting up the dialer
	d := gomail.NewDialer("criptbond@gmail.com", 587, "", "9GSpLZ7hdgRUmJnN")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return true
}
