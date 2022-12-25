package Tickets

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// most important
var CollectionMongo *mongo.Collection
var testVariable = "Sidharth"

// connect with mongodb
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var connectionString = os.Getenv("MONGO_URL")
	const dbName = "netflix"
	const colName = "Tickets"

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	CollectionMongo = client.Database(dbName).Collection(colName)
	fmt.Println("COllection instance is ready")
}
