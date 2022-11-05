package Auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
)
func LoginController(w http.ResponseWriter, r *http.Request) {
	//getting the emeail and password from the request body
	w.Header().Set("Content-Type", "application/json")
	//getting the emeial and passwprd from the mux
	params := r.URL.Query()
	var auth models.Auth = Login(params.Get("email"), params.Get("password"));
	Login(params.Get("email"), params.Get("password"))
	fmt.Println("Login controller is called");
	//creating the jwt token	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commons.Response{Status: 200,Message: "User is logged in", Data: auth})
}
//signup controller
func SignupController(w http.ResponseWriter, r *http.Request) {
	//getting the emeail and password from the request body
	w.Header().Set("Content-Type", "application/json")
	//getting the emeial and passwprd from the mux in request body
	var user models.Auth
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	Signup(user);
	fmt.Println("Signup controller is called");
	//creating the jwt token	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commons.Response{Status: 200,Message: "User is logged in"})
}
//test route 
func TestController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commons.Response{Status: 200,Message: "Test route is called"})
}