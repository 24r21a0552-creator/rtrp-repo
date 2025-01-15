package middleware

import (
	"context"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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

func EmailConfirmation(reciever, response, operation string) error {

	from := "hiimindra12345@gmail.com"
	password := "jtnu dbbl cqvn xetv"
	to := []string{reciever}
	auth := smtp.PlainAuth("smtp", from, password, "smtp.gmail.com")
	var subject, body string
	if operation == "booking" {
		subject = "Subject: Booking Confirmation\n"
		body = "Your booking has been confirmed. The following are the details:\n" + response

	}
	if operation == "cancelling" {
		subject = "Subject: Cancel Confirmation\n"
		body = "Your booking has been Cancelled. Details:\n" + response
	}

	message := []byte(subject + "\n" + body)
	smtpAddress := "smtp.gmail.com"
	smtpPort := "587"
	err := smtp.SendMail(smtpAddress+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil

}
