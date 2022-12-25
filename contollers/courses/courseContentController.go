package courses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
)

func CreateCourseStringContentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	var response commons.Response
	//getting the response in the form of request body
	var courseContent models.CourseContent
	_ = json.NewDecoder(r.Body).Decode(&courseContent)
	response = CreateStringContent(courseContent, courseId)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No Course Found"})
	}
}

// creating the function for the course File intake controller
// handler to handle the image upload
func UploadImages(w http.ResponseWriter, r *http.Request) {
	// Parse request body as multipart form data with 32MB max memory
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
	}

	// Get file from Form
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Create file locally
	dst, err := os.Create(handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer dst.Close()

	// Copy the uploaded file data to the newly created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)

}
