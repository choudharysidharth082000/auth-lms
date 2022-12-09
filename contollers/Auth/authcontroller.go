package Auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login controller is called")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	//getting the request body
	var auth models.Auth
	_ = json.NewDecoder(r.Body).Decode(&auth)
	userEmail := auth.Email
	userPassword := auth.Password
	//calling the login function in the methods
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Login(userEmail, userPassword))
}

func SignupController(w http.ResponseWriter, r *http.Request) {
	//getting the emeail and password from the request body
	
	fmt.Println("Signup controller is called")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// getting the emeial and passwprd from the mux in request body
	var user models.Auth
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	//status accepred
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Signup(user))
}
// test route
func TestController(w http.ResponseWriter, r *http.Request) {
	//cors handling
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	fmt.Println("Hwllo wolr vnrjvnr")
	//adding the default headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	var auth models.Auth
	_ = json.NewDecoder(r.Body).Decode(&auth)
	fmt.Println(auth.Email, auth.Password)
	// Login(auth.Email, auth.Password);
	//Aloow all headers
	// w.Header().Set("Allow-Control-Allow_Headers", "content-type");
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "Test route is called"})
}

// forhgot password route
func ForgotPasswordController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	if ForgetPassword(params.Get("email")) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "Email is sent to the user"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "Email is not sent to the user"})
	}
}

// verify otp
func VerifyOTPController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	if VerifyOTP(params.Get("email"), params.Get("otp")) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "OTP is verified"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "OTP is not verified"})
	}
}
