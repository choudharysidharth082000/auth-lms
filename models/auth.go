package models

//importing the primitive package from mongodb driver
import (
	// "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	UserName     string             `json:"username,omitempty" bson:"username,omitempty"`
	FirstName    string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName     string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Phone        string             `json:"phone,omitempty" bson:"phone,omitempty"`
	ProfileImage string             `json:"profileimage" bson:"profileimage"`
	Role         string             `json:"role,omitempty" bson:"role,omitempty"`
	Wishlist     []string           `json:"wishlist" bson:"wishlist"`
	Purchased    []string           `json:"purchased" bson:"purchased"`
	// jwt.StandardClaims
}
