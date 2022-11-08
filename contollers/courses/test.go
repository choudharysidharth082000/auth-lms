package courses

import (

	"github.com/sidharthchoudhary/lmsAuth/Mailer"
)
func sendEmailTest(text string, email string, body string) bool{
	sendEmail := mailer.MailUser(email, text, body);
	return sendEmail;
}