package Auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
	jwt "github.com/sidharthchoudhary/lmsAuth/utils/JWT"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	//getting the emeail and password from the request body
	w.Header().Set("Content-Type", "application/json")
	//getting the emeial and passwprd from the mux
	params := r.URL.Query()
	var auth models.Auth = Login(params.Get("email"), params.Get("password"))
	Login(params.Get("email"), params.Get("password"))
	fmt.Println("Login controller is called")
	//creating the jwt token
	w.WriteHeader(http.StatusOK)
	tokenGenerated, err := jwt.CreateJWT(auth.ID.Hex())
	if err != nil {
		json.NewEncoder(w).Encode(commons.Response{Status: 500, Message: "Error in creating the token"})
	}
	//creating the jwt token
	json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "User is logged in", Data: auth, Token: tokenGenerated})
}

// signup controller
func SignupController(w http.ResponseWriter, r *http.Request) {
	//getting the emeail and password from the request body
	w.Header().Set("Content-Type", "application/json")
	//getting the emeial and passwprd from the mux in request body
	var user models.Auth
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	Signup(user)
	fmt.Println("Signup controller is called")
	//creating the jwt token
	w.WriteHeader(http.StatusOK)
	createToken, err := jwt.CreateJWT(user.ID.Hex())
	if err != nil {
		json.NewEncoder(w).Encode(commons.Response{Status: 500, Message: "Error in creating the token"})
	}
	json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "User is logged in", Token: createToken})
}

// test route
func TestController(w http.ResponseWriter, r *http.Request) {
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
