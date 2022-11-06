package Auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
)

//controllers

func UpdateProfileController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//getting the request body
	var user models.Auth;
	json.NewDecoder(r.Body).Decode(&user);
	//calling te update profile function
	if UpdateProfile(user, params["id"]) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.Response{Status: 200,Message: "Profile is updated"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400,Message: "Profile is not updated"})
	}
}

//getting all the profiles from the user
func GetAllProfilesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	profiles := GetAllProfiles()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commons.Response{Status:200,Message:"All the profiles",Data:profiles});
	fmt.Println("Profile Router is Called");
}
//getting the profile by id
func GetProfileController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	profile := GetProfile(params["id"])
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commons.Response{Status:200,Message:"Profile",Data:profile});
	fmt.Println("Profile Router is Called");
}
