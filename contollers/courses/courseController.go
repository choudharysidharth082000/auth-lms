package courses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
)

func AddProductWish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Println(userId)
	productId := vars["courseId"]
	var response commons.Response
	response = AddWishListController(userId, productId)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "Product is added to wishlist"})
	} else {
		fmt.Println(response.Message)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "Product is not added to wishlist"})
	}
}

// adding the product to the cart
func AddProductCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	productId := vars["productId"]
	var response commons.Response
	response = AddProductToCartController(userId, productId)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.Response{Status: 200, Message: "Product is added to cart"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "Product is not added to cart"})
	}
}

// getting all the courses
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response commons.Response
	response = GetAllCoursesController()
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No courses found"})
	}
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var courseVal models.Course
	_ = json.NewDecoder(r.Body).Decode(&courseVal)
	var response commons.Response
	response = InsertCourse(courseVal)
	json.NewEncoder(w).Encode(response)
}

// controller to update the course
func UpdateCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var courseVal models.Course
	_ = json.NewDecoder(r.Body).Decode(&courseVal)
	//converting the course id to string
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	var response commons.Response
	response = UpdateCourse(courseVal, courseId)
	json.NewEncoder(w).Encode(response)
}

//controller for getting the single course

func GetSingleCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	var response commons.Response
	response = GetSingleCourse(courseId)
	json.NewEncoder(w).Encode(response)
}

//controller for deleting the course

func DeleteCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	var response commons.Response
	response = DeleteCourseFunction(courseId)
	json.NewEncoder(w).Encode(response)
}
//deleting all the courses by the course id
func DeleteAllCoursesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response commons.Response
	//getting teh course id from teh url
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	response = DeleteManyCourses(courseId)
	json.NewEncoder(w).Encode(response)

}
//deleting all the courses by the user id
func DeleteAllCoursesByUserIdController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response commons.Response
	//getting teh course id from teh url
	vars := mux.Vars(r)
	userId := vars["userId"]
	response = DeleteManyCoursesByUser(userId)
	json.NewEncoder(w).Encode(response)

}