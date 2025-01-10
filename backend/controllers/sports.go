package controllers

import (
	"net/http"
	"os"
	"sportslotbooker/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type Controllers interface {
	CreateBooking(w http.ResponseWriter, r *http.Request)
	CancelBooking(w http.ResponseWriter, r *http.Request)
}

type SportsControllers struct {
	service services.Services
}

func (s *SportsControllers) CreateBooking(w http.ResponseWriter, r *http.Request) {}
func (s *SportsControllers) CancelBooking(w http.ResponseWriter, r *http.Request) {}

func NewController(client *mongo.Client) Controllers {
	db := client.Database(os.Getenv("DATABASE"))
	coll := db.Collection(os.Getenv("BOOKINGS"))
	ser := services.NewService(coll)
	return &SportsControllers{
		service: ser,
	}
}
