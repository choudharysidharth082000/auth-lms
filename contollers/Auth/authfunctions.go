package Auth

import (
	"context"
	"fmt"
	"log"

	"github.com/pquerna/otp/totp"
	mailer "github.com/sidharthchoudhary/lmsAuth/Mailer"
	"github.com/sidharthchoudhary/lmsAuth/models"
	validate "github.com/sidharthchoudhary/lmsAuth/utils/Validate"
	"go.mongodb.org/mongo-driver/bson"
)

//importing the jwt package

func Login(email string, password string) models.Auth {
	//is the email 8 characaters long
	if !validate.LoginValidate(email, password) {
		log.Fatal("Email or password is not valid")
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
	if !validate.SignupValidate(user.Email, user.Password, user.FirstName, user.LastName, user.UserName, user.Phone) {
		log.Fatal("Email or password is not valid")
	}
	//creating a variable for the auth
	var auth models.Auth
	filter := bson.M{"email": user.Email}
	//checking for the email in the database
	err := collection.FindOne(context.TODO(), filter).Decode(&auth)
	if err == nil {
		log.Fatal("User Already present")
		log.Fatal(err)
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

// function to generate otp
func GenerateOTP(email string) string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	//sending the otp to the email
	sendMail := mailer.MailUser(email, "One Time Password", key.Secret())
	if sendMail {
		return key.Secret()
	}
	return "Not Found"
}

// forget password function
func ForgetPassword(email string) bool {
	//finding the user
	var auth models.Auth
	filter := bson.M{"email": email}
	//checking for the email in the database
	err := collection.FindOne(context.TODO(), filter).Decode(&auth)
	if err != nil {
		log.Fatal(err)
	}
	//generating the otp
	otp, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "Sidharth",
	})
	if err != nil {
		log.Fatal(err)
	}
	//updating the database
	updateRecords, err := collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", bson.D{
			{"otp", otp.Secret()},
		}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateRecords.MatchedCount, updateRecords.ModifiedCount)
	//sending the otp to the email
	sendMail := mailer.MailUser(email, "One Time Password", otp.Secret())
	if sendMail {
		return true
	}
	return false
}
//verify otp function
func VerifyOTP(email string, otp string) bool {
	var auth models.Auth
	filter := bson.M{"email": email}
	//checking for the email in the database
	err := collection.FindOne(context.TODO(), filter).Decode(&auth)
	if err != nil {
		log.Fatal(err)
	}
	//checking if the otp is correct
	if auth.OTP.OTP == otp {
		return true
	}
	return false
}