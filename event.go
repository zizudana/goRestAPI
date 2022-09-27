package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type eventStruct struct {
	CarName string             `json:"title" bson:"title"`
	Users   string             `json:"display" bson:"display"`
	Start   primitive.DateTime `json:"start" bson:"start"`
	End     primitive.DateTime `json:"end" bson:"end"`
}

type eventStructWithObjectID struct {
	ObjectID primitive.ObjectID `json:"_id" bson:"_id"`
	ID       primitive.ObjectID `json:"id" bson:"id"`
	CarName  string             `json:"title" bson:"title"`
	Users    string             `json:"display" bson:"display"`
	Start    primitive.DateTime `json:"start" bson:"start"`
	End      primitive.DateTime `json:"end" bson:"end"`
}

func initEventContent(e *echo.Echo) {
	e.POST("/events", createEvent)
	e.GET("/events/all", readAllEvent)
	e.DELETE("/events", deleteEvent)
}

func createEvent(c echo.Context) error {
	newEvent := new(eventStruct)

	err := c.Bind(newEvent)
	errCheck(err)

	insertOneResult, err := collection["car_event"].InsertOne(
		ctx,
		newEvent,
	)
	errCheck(err)

	fmt.Println(newEvent)

	return c.JSON(http.StatusCreated, insertOneResult)
}

func readAllEvent(c echo.Context) error {
	cur, err := collection["car_event"].Find(
		ctx,
		bson.M{},
	)
	fmt.Println(cur)
	errCheck(err)
	defer cur.Close(ctx)

	eventArr := []*eventStruct{}

	for cur.Next(ctx) {
		eventResult := new(eventStruct)
		err := cur.Decode(&eventResult)
		errCheck(err)

		eventArr = append(eventArr, eventResult)
	}
	if err := cur.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{
			"message": "FAIL readWrongContentAll",
		})
	}

	//logger.Info("SUCCESS readWrongContentAll")

	response := bson.M{
		"event_arr": eventArr,
	}
	fmt.Println(response)
	return c.JSON(http.StatusOK, response)
}

func deleteEvent(c echo.Context) error {
	hex := c.QueryParam("objectid")
	objectID, err := primitive.ObjectIDFromHex(hex)
	errCheck(err)

	deleteResult, err := collection["event"].DeleteOne(
		ctx,
		bson.M{
			"_id": objectID,
		},
	)
	errCheck(err)

	//logger.Info("SUCCESS deleteVideo : %s", hex)

	return c.JSON(http.StatusOK, deleteResult)
}
