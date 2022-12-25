package Tickets

import (
	"context"
	"time"

	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	"github.com/sidharthchoudhary/lmsAuth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTicketFunction(ticket models.Ticket, courseID string, userID string) commons.Response {
	//creating a new ticket
	//cehcking if the user id existst or not
	var auth models.Auth
	filter := bson.M{"_id": userID}
	err := CollectionMongo.FindOne(context.TODO(), filter).Decode(&auth)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "User Not Found",
		}
	}

	ticket.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	ticket.IsAnswered = false;
	ticket.CourseId = courseID;
	ticket.UserId = userID;
	ticket.UserName = auth.UserName;
	ticket.IsAnswered = false
	insertedData, err := CollectionMongo.InsertOne(context.TODO(), ticket)
	if err != nil {
		return commons.Response{Status: 400, Message: "Error in creating ticket"}
	}
	return commons.Response{Status: 1, Message: "Ticket Created", Data: insertedData}
}

// get all tickets
func GetAllTicketsFunction() commons.Response {
	var tickets []models.Ticket
	cursor, err := CollectionMongo.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return commons.Response{
			Status:  400,
			Message: "Error in getting tickets",
		}
	}
	//cursor loop
	for cursor.Next(context.TODO()) {
		var ticket models.Ticket
		err := cursor.Decode(&ticket)
		if err != nil {
			return commons.Response{
				Status:  400,
				Message: "Error in getting tickets",
			}
		}
		tickets = append(tickets, ticket)
	}
	if err := cursor.Err(); err != nil {
		return commons.Response{
			Status:  400,
			Message: "Error in getting tickets",
		}
	}
	return commons.Response{Status: 1, Message: "Tickets", Data: tickets}
}

// get tickets by ticketID
func GetTicketByIDFunction(ticketID string) commons.Response {
	var ticket models.Ticket
	ticketURL, err := primitive.ObjectIDFromHex(ticketID)
	if err != nil {
		return commons.Response{Status: 400, Message: "Error in getting ticket"}
	}
	filter := bson.M{"_id": ticketURL}
	err = CollectionMongo.FindOne(context.TODO(), filter).Decode(&ticket)
	if err != nil {
		return commons.Response{Status: 400, Message: "Error in getting ticket"}
	}
	return commons.Response{Status: 1, Message: "Ticket", Data: ticket}
}

// get tucket by userID
func GetTicketsByUserIDFunction(userID string) commons.Response {
	var tickets []models.Ticket
	filter := bson.M{"userId": userID}
	cursor, err := CollectionMongo.Find(context.TODO(), filter)
	if err != nil {
		return commons.Response{Status: 400, Message: "Error in getting tickets"}
	}
	for cursor.Next(context.TODO()) {
		var ticket models.Ticket
		cursor.Decode(&ticket)
		tickets = append(tickets, ticket)
	}
	return commons.Response{Status: 1, Message: "Tickets", Data: tickets}
}

// get tickets by courseID
func GetTicketsByCourseIDFunction(courseID string, userID string) commons.Response {
	var tickets []models.Ticket
	cursor, err := CollectionMongo.Find(context.TODO(), models.Ticket{CourseId: courseID})
	if err != nil {
		return commons.Response{Status: 400, Message: "Error in getting tickets"}
	}
	for cursor.Next(context.TODO()) {
		var ticket models.Ticket
		cursor.Decode(&ticket)
		tickets = append(tickets, ticket)
	}
	return commons.Response{Status: 1, Message: "Tickets", Data: tickets}
}

// answer the ticket
func AnswerTicketFunction(courseID string, userID string, ticketID string, answer string) commons.Response {
	ticketURL, err := primitive.ObjectIDFromHex(ticketID)
	if err != nil {
		return commons.Response{Status: 400, Message: "Error in getting ticket"}
	}
	filter := bson.D{{"courseId", courseID}, {"_id", ticketURL}}
	update := bson.D{
		{"$set", bson.D{
			{"answer", answer},
			{"isAnswered", true},
		}},
	}
	_, err = CollectionMongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return commons.Response{Status: 400, Message: "Error in updating ticket"}
	}
	return commons.Response{Status: 1, Message: "Ticket Updated"}
}
