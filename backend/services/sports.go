package services

import (
	"context"
	"log"
	"sportslotbooker/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services interface {
	Create(booking model.Booking) error
	Cancel(cancel model.Cancellation) error
}

type SportsServices struct {
	collection *mongo.Collection
}

func NewService(coll *mongo.Collection) Services {
	return &SportsServices{
		collection: coll,
	}
}

func (c *SportsServices) Create(booking model.Booking) error {

	_, err := c.collection.InsertOne(context.TODO(), booking)

	if err != nil {
		log.Println("error during insertion into the db ")
		return err
	}
	// send e-mail
	return nil

}
func (c *SportsServices) Cancel(cancel model.Cancellation) error {
	// Use the filter variable in a MongoDB operation to avoid "declared and not used" error
	roll, date, sport := cancel.Roll_no, cancel.Date, cancel.Sport
	filter := bson.D{
		{Key: "roll_no", Value: roll}, {Key: "date", Value: date}, {Key: "sport", Value: sport}}
	res, err := c.collection.DeleteOne(context.Background(), filter)
	if err != nil && res.DeletedCount == 1 {
		log.Println("error during deletion from the db")
		return err
	}
	return nil
}
