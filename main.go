package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var logger = helper.NewColorLogger()
var myLogger *log.Logger
var ctx context.Context

var emptyObjectID, _ = primitive.ObjectIDFromHex("0000000000000000")

func main() {
	myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	e := echo.New()

	setEchoMiddleware(e)

	ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancle()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://zizudana:zizudana32@carreserve.fk72gpm.mongodb.net/?retryWrites=true&w=majority"))
	defer disconnectMongo(client)
	errCheck(err)

	myLogger.Println("SUCCESS connect mongo")

	setCollection(client)

	// Static files
	e.Static("/files", "files")

	// Routes
	initEventContent(e)

	// Start server
	//e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	e.Logger.Fatal(e.Start(":1323"))
}
