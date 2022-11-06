package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sidharthchoudhary/lmsAuth/contollers/Auth"
)

func main() {
	newRouter := mux.NewRouter()
	newRouter.HandleFunc("/v1/api/auth/login", Auth.LoginController).Methods("GET")
	newRouter.HandleFunc("/v1/api/auth/signup", Auth.SignupController).Methods("POST")
	newRouter.HandleFunc("/test", Auth.TestController).Methods("GET")

	//profile routes
	newRouter.HandleFunc("/v1/api/profile/getProfile/{id}", Auth.GetProfileController).Methods("GET")
	newRouter.HandleFunc("/v1/api/profile/updateProfile", Auth.UpdateProfileController).Methods("PUT")
	newRouter.HandleFunc("/v1/api/profile/getProfile", Auth.GetAllProfilesController).Methods("GET")
	//serve the server on port 4040
	fmt.Println("Server is running on port 4040")
	log.Fatal(http.ListenAndServe(":4040", newRouter))	

}
