package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sidharthchoudhary/lmsAuth/contollers/Auth"
	"github.com/sidharthchoudhary/lmsAuth/contollers/courses"
)

func main() {
	newRouter := mux.NewRouter()
	newRouter.HandleFunc("/v1/api/auth/login", Auth.LoginController).Methods("POST", "OPTIONS")
	newRouter.HandleFunc("/v1/api/auth/signup", Auth.SignupController).Methods("POST", "OPTIONS")
	newRouter.HandleFunc("/test", Auth.TestController).Methods("POST", "OPTIONS")
	newRouter.HandleFunc("/v1/api/auth/forgotPassword", Auth.ForgotPasswordController).Methods("POST", "OPTIONS")
	//profile routes
	newRouter.HandleFunc("/v1/api/profile/getProfile/{id}", Auth.GetProfileController).Methods("GET")
	newRouter.HandleFunc("/v1/api/profile/updateProfile", Auth.UpdateProfileController).Methods("PUT")
	newRouter.HandleFunc("/v1/api/profile/getProfile", Auth.GetAllProfilesController).Methods("GET")
	//courses routes
	newRouter.HandleFunc("/v1/api/profile/addWishList/{courseId}/{userId}", Auth.AddProductWish).Methods("POST")
	newRouter.HandleFunc("/v1/api/courses/getCourse/{courseId}/{userId}", Auth.AddProductCart).Methods("POST")
	newRouter.HandleFunc("/v1/api/courses/getAllCourse", courses.GetAllCourses).Methods("GET")
	newRouter.HandleFunc("/v1/api/courses/createCourse", courses.CreateCourse).Methods("POST")
	newRouter.HandleFunc("/v1/api/courses/updateCourse/{courseId}", courses.UpdateCourseController).Methods("PUT")
	newRouter.HandleFunc("/v1/api/courses/getAllCourse/{courseId}", courses.GetSingleCourseController).Methods("GET")
	newRouter.HandleFunc("/v1/api/courses/deleteCourse/{courseId}", courses.DeleteCourseController).Methods("DELETE")
	newRouter.HandleFunc("/v1/api/courses/deleteCourseByCourseID/{courseId}", courses.DeleteAllCoursesController).Methods("DELETE")
	newRouter.HandleFunc("/v1/api/courses/deleteCourseByUserID/{userId}", courses.DeleteAllCoursesByUserIdController).Methods("DELETE")

	//course content routers
	newRouter.HandleFunc("/v1/api/couses/addCourseContent/{courseId}", courses.CreateCourseStringContentController).Methods("POST")
	//test route for sending email
	newRouter.HandleFunc("/v1/api/test", courses.TestController).Methods("GET")
	//serve the server on port 4040
	fmt.Println("Server is running on port 4040")
	log.Fatal(http.ListenAndServe(":4040", newRouter))

}
