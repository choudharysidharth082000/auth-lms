package Auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pquerna/otp/totp"
	commons "github.com/sidharthchoudhary/lmsAuth/Commons"
	mailer "github.com/sidharthchoudhary/lmsAuth/Mailer"
	"github.com/sidharthchoudhary/lmsAuth/models"
	jwt "github.com/sidharthchoudhary/lmsAuth/utils/JWT"
	validate "github.com/sidharthchoudhary/lmsAuth/utils/Validate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//importing the jwt package

func Login(email string, password string) commons.Response {
	if validate.LoginValidate(email, password) == false {
		return commons.Response{
			Status:  0,
			Message: "Invalid Email or Password",
		}
	}
	var auth models.Auth
	filter := bson.M{"email": email}
	err := CollectionMongo.FindOne(context.TODO(), filter).Decode(&auth)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "User Not Found",
		}
	}
	generatedToken, err := jwt.CreateJWT(auth.ID.Hex())
	fmt.Println(auth.Password)
	if auth.Password == password {
		return commons.Response{
			Status:  1,
			Message: "Login Successful",
			Data:    auth,
			Token: generatedToken,
		}
	}
	fmt.Println(auth)
	return commons.Response{
		Status:  0,
		Message: "Invalid Email or Password",
	}

}


func Signup(user models.Auth) commons.Response {
	//is the email 8 characaters long
	fmt.Println(user.Email)
	if !validate.SignupValidate(user.Email, user.Password, user.UserName) {
		return commons.Response{
			Status:  0,
			Message: "Invalid Email or Password",
		}
	}
	//creating a variable for the auth
	var auth models.Auth
	filter := bson.M{"email": user.Email}
	//checking for the email in the database
	err := CollectionMongo.FindOne(context.TODO(), filter).Decode(&auth)
	if err == nil {
		return commons.Response{
			Status:  0,
			Message: "User Already Exists",
			Data:    nil,
			Token:   "",
		}
	}
	//inserting the user in the database
	insertedData, err := CollectionMongo.InsertOne(context.TODO(), user)
	if err != nil {
		return commons.Response{
			Status:  0,
			Message: "Error Creating User",
		}
	}
	//creating the token from the resoonse
	generatedToken, err := jwt.CreateJWT(insertedData.InsertedID.(primitive.ObjectID).Hex())
	return commons.Response{
		Status:  1,
		Message: "User Created Successfully",
		Data:    insertedData,
		Token:   generatedToken,
	}
}

// getting all the data
// getting all records from mongodb
func getAllMovies() []models.Auth {
	var movies []models.Auth
	cur, err := CollectionMongo.Find(context.TODO(), bson.D{{}})
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
	err := CollectionMongo.FindOne(context.TODO(), filter).Decode(&auth)
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
	updateRecords, err := CollectionMongo.UpdateOne(context.TODO(), filter, bson.D{
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

// verify otp function
func VerifyOTP(email string, otp string) bool {
	var auth models.Auth
	filter := bson.M{"email": email}
	err := CollectionMongo.FindOne(context.TODO(), filter).Decode(&auth)
	if err != nil {
		fmt.Println(err)
		return false
	}
	//calculating the time difference
	timeDifference := time.Now().UnixNano() - int64(auth.OTP.CreatedAt)
	//checking the time difference
	if timeDifference > 30000000000 {
		return false
	}
	//checking the otp
	if auth.OTP.OTP == otp {
		//updating the otp in the database
		var varOTP models.OTP
		varOTP.Email = email
		varOTP.OTP = otp
		varOTP.Verified = true
		varOTP.CreatedAt = primitive.DateTime(time.Now().UnixNano())
		//updating the data to the database
		insertedData, err := CollectionMongo.UpdateOne(context.TODO(), models.OTP{Email: email}, varOTP)
		if err != nil {
			fmt.Println(err)
			return false
		}
		fmt.Println(insertedData)
		return true
	} else {
		return false
	}
}
