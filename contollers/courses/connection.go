package courses

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sidharth:sidharth@cluster0.nlhjv1y.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "Course"

// most important
var CollectionMongo *mongo.Collection
var testVariable = "Sidharth"

// connect with mongodb
func init() {
	fmt.Println("Course Collection Creation Called fucntion...");
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	CollectionMongo = client.Database(dbName).Collection(colName)
	fmt.Println("COllection instance is ready")
}
