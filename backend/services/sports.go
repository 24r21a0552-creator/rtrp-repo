package services

import (
	"context"
	"log"
	"sportslotbooker/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Services interface {
	Create(booking model.Booking) error
	Cancel() error
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
func (c *SportsServices) Cancel() error {}
