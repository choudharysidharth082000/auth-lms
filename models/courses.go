package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	Id             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID         string             `json:"userid,omitempty" bson:"userid,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	Image          string             `json:"image,omitempty" bson:"image,omitempty"`
	BannerImage    string             `json:"bannerImage,omitempty" bson:"bannerImage"`
	Price          int                `json:"price,omitempty" bson:"price,omitempty"`
	CreatedBy      string             `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	UpdatedAt      string             `json:"updatedAt,omitempty" bson	:"updatedAt,omitempty"`
	CreatedAt      string             `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	CourseContents []CourseContent    `json:"courseContents" bson:"courseContents"`
}

type CourseContent struct {
	Id                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ContentName        string             `json:"contentName,omitempty" bson:"contentName,omitempty"`
	ContentDescription string             `json:"contentDescription,omitempty" bson:"contentDescription,omitempty"`
	ContentImage       string             `json:"contentImage,omitempty" bson:"contentImage,omitempty"`
	ContentVideo       string             `json:"contentVideo,omitempty" bson:"contentVideo,omitempty"`
	ContentAudio       string             `json:"contentAudio,omitempty" bson:"contentAudio,omitempty"`
	ContentFile        string             `json:"contentFile,omitempty" bson:"contentFile,omitempty"`
	ContentLink        string             `json:"contentLink,omitempty" bson:"contentLink,omitempty"`
	CreatedAt          string             `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt          string             `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
