package Auth

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sidharth:sidharth@cluster0.nwnkaln.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "Auth"

// most important
var collection *mongo.Collection
var testVariable = "Sidharth"

// connect with mongodb
func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("COllection instance is ready")
}
