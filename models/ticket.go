package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CourseId   string             `json:"courseid,omitempty" bson:"courseid,omitempty"`
	UserId     string             `json:"userid,omitempty" bson:"userid,omitempty"`
	UserName   string             `json:"username,omitempty" bson:"username,omitempty"`
	Topic      string             `json:"topic,omitempty" bson:"topic,omitempty"`
	Question   string             `json:"question,omitempty" bson:"question,omitempty"`
	Answer     string             `json:"answer" bson:"answer"`
	IsAnswered bool               `json:"isanswered,omitempty" bson:"isanswered,omitempty"`
	AnsweredBy string             `json:"answeredby,omitempty" bson:"answeredby,omitempty"`
	AnsweredAt primitive.DateTime `json:"answeredat,omitempty" bson:"answeredat,omitempty"`
	CreatedAt  primitive.DateTime `json:"createdat,omitempty" bson:"createdat,omitempty"`

	// jwt.StandardClaims
}
