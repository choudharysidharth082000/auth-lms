package courses

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sidharthchoudhary/lmsAuth/models"
	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
)

func CreateCourseStringContentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	var response commons.Response
	//getting the response in the form of request body
	var courseContent models.CourseContent
	_ = json.NewDecoder(r.Body).Decode(&courseContent)
	response = CreateStringContent(courseContent,courseId)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No Course Found"})
	}
}
