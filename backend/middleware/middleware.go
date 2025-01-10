package middleware

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo"
)

func DBConnection() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err in loading the env ", err)
		return nil
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("error during connection", err)
		return nil
	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("error during the Ping", err)
	}
	log.Println("Connection Established")
	return mongoClient
}
