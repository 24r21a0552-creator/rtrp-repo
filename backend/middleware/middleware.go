package middleware

import (
	"context"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func EmailConfirmation(reciever, response string) error {

	from := "nbkreddy12345@gmail.com"
	password := "6302503241"
	message := []byte(response)
	to := []string{reciever}
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	smtpAddress := "smtp.gmail.com"
	smtpPort := "587"
	err := smtp.SendMail(smtpAddress+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil

}
