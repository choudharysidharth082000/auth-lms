package Tickets

import (
	"encoding/json"
	"fmt"
	"net/http"

	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	//getting the discription from the email and userid and courseId from the parameters
	fmt.Println("Controller for creat ticket is called");
	params := r.URL.Query()
	userID := params.Get("userID")
	courseID := params.Get("courseID")
	//userID
	//courseID
	var ticket models.Ticket
	_ = json.NewDecoder(r.Body).Decode(&ticket)
	fmt.Println(ticket)
	response := CreateTicketFunction(ticket, courseID, userID)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No Course Found"})
	}
}

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	//getting the discription from the email and userid and courseId from the parameters
	//userID
	//courseID
	response := GetAllTicketsFunction()
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No Course Found"})
	}
}

// getting tickets by ticketID
func GetTicketByID(w http.ResponseWriter, r *http.Request) {
	//getting the discription from the email and userid and courseId from the parameters
	params := r.URL.Query()
	ticketID := params.Get("ticketID")
	//userID
	//courseID
	response := GetTicketByIDFunction(ticketID)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No Course Found"})
	}
}

// getting tickets by courseID
func GetTicketsByCourseID(w http.ResponseWriter, r *http.Request) {
	//getting the discription from the email and userid and courseId from the parameters
	params := r.URL.Query()
	userID := params.Get("userID")
	courseID := params.Get("courseID")
	//userID
	//courseID
	response := GetTicketsByCourseIDFunction(courseID, userID)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No Course Found"})
	}
}

// getting tickets by userID
func GetTicketsByUserID(w http.ResponseWriter, r *http.Request) {
	//getting the discription from the email and userid and courseId from the parameters
	params := r.URL.Query()
	userID := params.Get("userID")
	//userID
	response := GetTicketsByUserIDFunction(userID)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No Course Found"})
	}
}

// answer the ticket
func AnswerTicket(w http.ResponseWriter, r *http.Request) {
	//getting the discription from the email and userid and courseId from the parameters
	params := r.URL.Query()
	ticketID := params.Get("ticketID")
	courseID := params.Get("courseID")
	userID := params.Get("userID")
	var ticket models.Ticket
	_ = json.NewDecoder(r.Body).Decode(&ticket)
	response := AnswerTicketFunction(ticketID, courseID, userID, ticket.Answer)
	if response.Status == 1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commons.Response{Status: 400, Message: "No Course Found"})
	}
}
