package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connec struct {
	Client *mongo.Client
}

func NewConnect(c *mongo.Client) *Connec {
	const uri = "mongodb://127.0.0.1:27017"
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
		return nil
	}
	fmt.Println("Successfully connected and pinged")
	return &Connec{
		Client: db,
	}

}
