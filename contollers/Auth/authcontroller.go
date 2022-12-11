package Auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pquerna/otp/totp"
	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	Mailer "github.com/sidharthchoudhary/lmsAuth/Mailer"
	"github.com/sidharthchoudhary/lmsAuth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// function for the forget password
func ForgotPasswordController(w http.ResponseWriter, r *http.Request) {
	//writing the required headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	//getting the email from the query
	params := r.URL.Query()
	email := params.Get("email")
	fmt.Println(email)
	//generating the otp and sending to the requested email
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "prodigalai.com",
		AccountName: "sidharth@prodigalai.com",
	})
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(commons.Response{
			Status:  0,
			Message: "Error in generating the otp",
		})
	}
	fmt.Println(key.Secret())
	var varOTP models.OTP
	varOTP.Email = email
	varOTP.OTP = key.Secret()
	varOTP.Verified = false
	//present time
	varOTP.CreatedAt = primitive.DateTime(time.Now().UnixNano())
	//inserting the data to the database
	insertedData, err := CollectionMongo.UpdateOne(context.TODO(), models.OTP{Email: email}, varOTP)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(commons.Response{
			Status:  0,
			Message: "Error in inserting the data to the database",
		})
	}
	fmt.Println(insertedData)
	//send the response to the user for success mail
	if Mailer.MailUser(email, "Secret for your Application", key.Secret()) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "OTP is sent to your email"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "OTP is not sent to your email"})
	}
}

// verify otp
func VerifyOTPController(w http.ResponseWriter, r *http.Request) {
	//headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	//getting the otp from the requested url
	params := r.URL.Query()
	otp := params.Get("otp")
	email := params.Get("email")
	//getting the otp from the data base
	var auth models.Auth
	filter := bson.M{"email": email}
	err := CollectionMongo.FindOne(context.TODO(), filter).Decode(&auth)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(commons.Response{Status: 0, Message: "Error in getting the data from the database"})
	}
	//calculating the time difference
	timeDifference := time.Now().UnixNano() - int64(auth.OTP.CreatedAt)
	//checking the time difference
	if timeDifference > 30000000000 {
		json.NewEncoder(w).Encode(commons.Response{Status: 0, Message: "OTP is expired"})
	}
	//checking the otp
	if auth.OTP.OTP == otp {
		//updating the otp in the database
		var varOTP models.OTP
		varOTP.Email = email
		varOTP.OTP = otp
		varOTP.Verified = true
		varOTP.CreatedAt = primitive.DateTime(time.Now().UnixNano())
		//updating the data to the database
		insertedData, err := CollectionMongo.UpdateOne(context.TODO(), models.OTP{Email: email}, varOTP)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(commons.Response{
				Status:  0,
				Message: "Error in inserting the data to the database",
			})
		}
		fmt.Println(insertedData)
		json.NewEncoder(w).Encode(commons.Response{Status: 1, Message: "OTP is verified"})
	} else {
		json.NewEncoder(w).Encode(commons.Response{Status: 0, Message: "OTP is not verified"})
	}
}
