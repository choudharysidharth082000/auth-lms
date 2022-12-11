package models

//importing the primitive package from mongodb driver
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	UserName     string             `json:"username,omitempty" bson:"username,omitempty"`
	ProfileImage string             `json:"profileimage" bson:"profileimage"`
	UserType     string             `json:"userType,omitempty" bson:"userType,omitempty"`
	Wishlist     []string           `json:"wishlist" bson:"wishlist"`
	Purchased    []string           `json:"purchased" bson:"purchased"`
	OTP          *OTP               `json:"otp" bson:"otp"`
	WishlistID   []string           `json:"wishlistid,eligible_for" bson:"wishlistid,eligible_for"`
	PurchasedID  []string           `json:"purchasedid" bson:"purchasedid"`
	// jwt.StandardClaims
}

type OTP struct {
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	OTP       string             `json:"otp,omitempty" bson:"otp,omitempty"`
	Verified  bool               `json:"verified,omitempty" bson:"verified,omitempty"`
	CreatedAt primitive.DateTime `json:"createdat,omitempty" bson:"createdat,omitempty"`
}

type Login struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}
