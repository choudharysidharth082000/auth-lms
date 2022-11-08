package courses

import (
	"encoding/json"
	"net/http"

	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	mailer "github.com/sidharthchoudhary/lmsAuth/Mailer"
)

func TestController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sendEmail := mailer.MailUser("choudharysidharth082000@gmail.com", "Test", "This is a test email")
	json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "Test route is called", Data: sendEmail})
}
