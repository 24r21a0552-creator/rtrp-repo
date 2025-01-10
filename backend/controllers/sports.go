package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sportslotbooker/model"
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

func (s *SportsControllers) CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	var booking model.Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("bad form ", booking)
		return
	}
	err = s.service.Create(booking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("created a new booking"))
	return

}
func (s *SportsControllers) CancelBooking(w http.ResponseWriter, r *http.Request) {}

func NewController(client *mongo.Client) Controllers {
	db := client.Database(os.Getenv("DATABASE"))
	coll := db.Collection(os.Getenv("BOOKINGS"))
	ser := services.NewService(coll)
	return &SportsControllers{
		service: ser,
	}
}
