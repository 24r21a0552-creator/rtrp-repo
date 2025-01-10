package services

import "go.mongodb.org/mongo-driver/mongo"

type Services interface {
	Create()
	Cancel()
}

type SportsServices struct {
	collection *mongo.Collection
}

func (c *SportsServices) Create() {}
func (c *SportsServices) Cancel() {}
func NewService(coll *mongo.Collection) Services {
	return &SportsServices{
		collection: coll,
	}
}
