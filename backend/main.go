package main

import (
	"net/http"
	"sportslotbooker/controllers"
	"sportslotbooker/middleware"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	connection := middleware.DBConnection()
	con := controllers.NewController(connection)
	router.HandleFunc("/CreateBooking", con.CreateBooking).Methods("POST")
	router.HandleFunc("/CancelBooking", con.CancelBooking).Methods("POST")

	http.ListenAndServe(":3000", router)
}
