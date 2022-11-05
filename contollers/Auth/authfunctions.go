package Auth

import (
	"context"
	"fmt"
	"log"

	"github.com/sidharthchoudhary/lmsAuth/models"
	"go.mongodb.org/mongo-driver/bson"
)

//importing the jwt package

func Login(email string, password string) models.Auth {
	//is the email 8 characaters long
	if len(email) < 8 {
		log.Fatal("Email is not 8 characters long")
	}
	//is the password 8 characters long
	if len(password) < 8 {
		log.Fatal("Password is not 8 characters long")
	}
	//creating a variable for the auth
	var auth models.Auth
	// convertId, _ := primitive.ObjectIDFromHex("63667e623638eddfb8f47c07")
	filter := bson.M{"email": email}
	fmt.Println(filter)
	//checking for the email in the database
	err := collection.FindOne(context.TODO(), filter).Decode(&auth)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	return auth
}

// signinig up
func Signup(user models.Auth) {
	//is the email 8 characaters long
	fmt.Println(user.Email)
	if len(user.Email) < 8 {
		log.Fatal("Email is not 8 characters long")
	}
	//is the password 8 characters long
	if len(user.Password) < 8 {
		log.Fatal("Password is not 8 characters long")
	}
	//creating a variable for the auth
	var auth models.Auth
	filter := bson.M{"email": user.Email}
	//checking for the email in the database
	err := collection.FindOne(context.TODO(), filter).Decode(&auth)
	if err == nil {
		log.Fatal("User does not exist")
	}
	//inserting the user in the database
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User is signed up")
}

// getting all the data
// getting all records from mongodb
func getAllMovies() []models.Auth {
	var movies []models.Auth
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var movie models.Auth
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return movies
}
