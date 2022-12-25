package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sidharthchoudhary/lmsAuth/contollers/Auth"
	"github.com/sidharthchoudhary/lmsAuth/contollers/Tickets"
	"github.com/sidharthchoudhary/lmsAuth/contollers/courses"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	newRouter := mux.NewRouter()
	newRouter.HandleFunc("/v1/api/auth/login", Auth.LoginController).Methods("POST", "OPTIONS")
	newRouter.HandleFunc("/v1/api/auth/signup", Auth.SignupController).Methods("POST", "OPTIONS")
	newRouter.HandleFunc("/test", Auth.TestController).Methods("POST", "OPTIONS")
	newRouter.HandleFunc("/v1/api/auth/forgotPassword", Auth.ForgotPasswordController).Methods("POST", "OPTIONS")
	//profile routes
	newRouter.HandleFunc("/v1/api/profile/getProfile/{id}", Auth.GetProfileController).Methods("GET")
	newRouter.HandleFunc("/v1/api/profile/updateProfile/{userID}", Auth.UpdateProfileController).Methods("PUT", "OPTIONS")
	newRouter.HandleFunc("/v1/api/profile/getProfile", Auth.GetAllProfilesController).Methods("GET")
	//courses routes
	newRouter.HandleFunc("/v1/api/profile/addWishList/{courseId}/{userId}", Auth.AddProductWish).Methods("POST")
	newRouter.HandleFunc("/v1/api/courses/getCourse/{courseId}/{userId}", Auth.AddProductCart).Methods("POST")
	newRouter.HandleFunc("/v1/api/courses/getAllCourse", courses.GetAllCourses).Methods("GET")
	newRouter.HandleFunc("/v1/api/courses/createCourse", courses.CreateCourse).Methods("POST", "OPTIONS")
	newRouter.HandleFunc("/v1/api/courses/updateCourse/{courseId}", courses.UpdateCourseController).Methods("PUT")
	newRouter.HandleFunc("/v1/api/courses/getAllCourse/{courseId}", courses.GetSingleCourseController).Methods("GET")
	newRouter.HandleFunc("/v1/api/courses/deleteCourse/{courseId}", courses.DeleteCourseController).Methods("DELETE")
	newRouter.HandleFunc("/v1/api/courses/deleteCourseByCourseID/{courseId}", courses.DeleteAllCoursesController).Methods("DELETE")
	newRouter.HandleFunc("/v1/api/courses/deleteCourseByUserID/{userId}", courses.DeleteAllCoursesByUserIdController).Methods("DELETE")

	//course content routers
	newRouter.HandleFunc("/v1/api/couses/addCourseContent/{courseId}", courses.CreateCourseStringContentController).Methods("POST")
	newRouter.HandleFunc("/v1/api/test/uploader", courses.UploadImages).Methods("POST")

	//routers for the test ticket
	newRouter.HandleFunc("/v1/api/ticket/createTicket/{courseID}/{userID}", Tickets.CreateTicket).Methods("POST", "OPTIONS");
	newRouter.HandleFunc("/v1/api/ticket/getAllTickets", Tickets.GetAllTickets).Methods("GET", "OPTIONS");
	newRouter.HandleFunc("/v1/api/test/ticket/{ticketID}", Tickets.GetTicketByID).Methods("GET", "OPTIONS");
	newRouter.HandleFunc("/v1/api/test/ticket/{courseID}", Tickets.GetTicketsByCourseID).Methods("PUT", "OPTIONS");
	//test route for sending email
	newRouter.HandleFunc("/v1/api/test", courses.TestController).Methods("GET")
	//serve the server on port 4040
	fmt.Println("Server is running on port 4040")
	fmt.Println(os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), newRouter))

}
