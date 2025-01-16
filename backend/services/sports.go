package services

import (
	"context"
	"log"
	"sportslotbooker/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services interface {
	Create(booking model.Booking) (bool, error)
	Cancel(cancel model.Cancellation) (bool, error)
}

type SportsServices struct {
	collection *mongo.Collection
}

func NewService(coll *mongo.Collection) Services {
	return &SportsServices{
		collection: coll,
	}
}

func (c *SportsServices) Create(booking model.Booking) (bool, error) {
	roll := booking.Roll_no
	filter := bson.D{{Key: "roll_no", Value: roll}}
	var existingBooking model.Booking
	err := c.collection.FindOne(context.TODO(), filter).Decode(&existingBooking)
	if err == nil {
		log.Println("Already a registration exists")
		return true, nil
	} else if err != mongo.ErrNoDocuments {
		log.Println("error during finding the document: ", err)
		return false, err
	}
	_, err = c.collection.InsertOne(context.TODO(), booking)
	if err != nil {
		log.Println("error during insertion into the db: ", err)
		return false, err
	}
	return false, nil

}
func (c *SportsServices) Cancel(cancel model.Cancellation) (bool, error) {
	// Use the filter variable in a MongoDB operation to avoid "declared and not used" error
	roll, date, sport := cancel.Roll_no, cancel.Date, cancel.Sport
	filter := bson.D{
		{Key: "roll_no", Value: roll}, {Key: "date", Value: date}, {Key: "sport", Value: sport}}
	var existingBooking model.Booking
	err := c.collection.FindOne(context.TODO(), filter).Decode(&existingBooking)
	if err == nil {
		log.Println("The Booking exists")
	} else if err == mongo.ErrNoDocuments {
		log.Println("no such booking is found ", err)
		return true, err
	} else {
		log.Println("error during finding the document: ", err)
		return false, err

	}
	res, err := c.collection.DeleteOne(context.Background(), filter)
	if err != nil && res.DeletedCount != 1 {
		log.Println("error during deletion from the db")
		return false, err
	}
	return false, nil
}
