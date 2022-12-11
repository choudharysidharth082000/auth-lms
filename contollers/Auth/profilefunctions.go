package Auth

import (
	"context"
	"fmt"
	"log"

	"github.com/sidharthchoudhary/lmsAuth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProfile(filter models.Auth, id string) bool {
	//coverting the string id to the primitive type
	convertedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	//updating the profile in the database
	_, err = CollectionMongo.UpdateOne(context.TODO(), bson.M{"_id": convertedId}, bson.M{"$set": filter})
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Updated the profile");
	return true
}

func GetProfile(id string) models.Auth {
	//converting the id
	convertedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	//creating a variable for the auth
	var auth models.Auth
	//checking for the email in the database
	err = CollectionMongo.FindOne(context.TODO(), bson.M{"_id": convertedId}).Decode(&auth)
	if err != nil {
		log.Fatal(err)
	}
	return auth
}
//getting all the profiles from the database
func GetAllProfiles() []models.Auth {
	//creating a variable for the auth
	var auth []models.Auth
	//checking for the email in the database
	cursor, err := CollectionMongo.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var auth1 models.Auth
		err = cursor.Decode(&auth1)
		if err != nil {
			log.Fatal(err)
		}
		auth = append(auth, auth1)
	}
	return auth
}