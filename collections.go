package main

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var collection map[string]*mongo.Collection

// setCollection : MongoDB collection 연결
func setCollection(client *mongo.Client) {
	collection = make(map[string]*mongo.Collection)

	dbCar := client.Database("car")
	collection["car_event"] = dbCar.Collection("car_event")
}
