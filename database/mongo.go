package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client
var AccommodationCollection *mongo.Collection

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
