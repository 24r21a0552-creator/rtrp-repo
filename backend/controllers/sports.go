package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sportslotbooker/middleware"
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

	type form struct {
		Roll_no    string `json:"roll_no,omitempty"`
		Email      string `json:"email,omitempty"`
		Department string `json:"department,omitempty"`
		Sport      string `json:"sport,omitempty"`
		Date       string `json:"date,omitempty"`
		Time       string `json:"time,omitempty"`
		Venue      string `json:"venue,omitempty"`
	}
	var booking form
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("bad form ", booking)
		return
	}
	newBooking := model.NewBooking(booking.Roll_no, booking.Email, booking.Department, booking.Sport, booking.Date, booking.Time, booking.Venue)
	fmt.Println("creating a new booking ", newBooking)
	err = s.service.Create(newBooking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("new booking created")
	confirmation := fmt.Sprintf("Roll Number: %s\nSport: %s\nDate: %s\nVenue: %s\n", booking.Roll_no, booking.Sport, booking.Date, booking.Venue)
	err = middleware.EmailConfirmation(booking.Email, confirmation,"booking")
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Println("email error ", err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("created a new booking"))
}
func (s *SportsControllers) CancelBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	var cancelling model.Cancellation
	err := json.NewDecoder(r.Body).Decode(&cancelling)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("bad form ", cancelling)
		return
	}
	fmt.Println("new cancellation ", cancelling)
	err = s.service.Cancel(cancelling)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	confirmation := fmt.Sprintf("Roll Number: %s\nSport: %s\nDate: %s\n", cancelling.Roll_no, cancelling.Sport, cancelling.Date)
	err = middleware.EmailConfirmation(cancelling.Email, confirmation,"cancelling")
	fmt.Println("completed cancellation")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("cancelled booking"))
}

func NewController(client *mongo.Client) Controllers {
	db := client.Database(os.Getenv("DATABASE"))
	coll := db.Collection(os.Getenv("BOOKINGS"))
	ser := services.NewService(coll)
	return &SportsControllers{
		service: ser,
	}
}
