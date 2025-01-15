package main

import (
	"net/http"
	"sportslotbooker/controllers"
	"sportslotbooker/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	connection := middleware.DBConnection()
	con := controllers.NewController(connection)
	router.HandleFunc("/CreateBooking", con.CreateBooking).Methods("POST")
	router.HandleFunc("/CancelBooking", con.CancelBooking).Methods("POST")

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Use the CORS middleware
	handler := c.Handler(router)

	http.ListenAndServe(":3000", handler)
}
