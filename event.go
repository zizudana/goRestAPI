package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type eventStruct struct {
	EventID string             `json:"id" bson:"id"`
	CarName string             `json:"title" bson:"title"`
	Display string             `json:"display" bson:"display"`
	Start   primitive.DateTime `json:"start" bson:"start"`
	End     primitive.DateTime `json:"end" bson:"end"`
}

func initEventContent(e *echo.Echo) {
	e.GET("/events/all", readAllEvent)
	e.POST("/events/add", createEvent)
	e.DELETE("/events/:id", deleteEvent)
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
	hex := c.Param("id")
	fmt.Println("SUCCESS deleteVideo : ", hex)

	//objectID, err := primitive.ObjectIDFromHex(hex)
	//errCheck(err)
	deleteResult, err := collection["car_event"].DeleteOne(
		ctx,
		bson.M{"id": hex},
	)
	errCheck(err)

	fmt.Println("SUCCESS deleteVideo : ", hex)

	return c.JSON(http.StatusOK, deleteResult)
}
