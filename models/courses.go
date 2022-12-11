package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	BannerImage string             `json:"bannerImage,omitempty" bson:"bannerImage"`
	Price       int                `json:"price,omitempty" bson:"price,omitempty"`
	CreatedBy   string             `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	UpdatedAt   string             `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	CreatedAt   string             `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}