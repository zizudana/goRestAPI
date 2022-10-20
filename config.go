package main

import (
	"context"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// setEchoMiddleware : set echo middleware
func setEchoMiddleware(e *echo.Echo) {

	e.Logger.SetLevel(log.ERROR)

	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "\033[92m${method}\t${uri}\t\033[94m${status}\t${remote_ip}\033[0m\n",
		}))

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"localhost"},
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}

func disconnectMongo(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	errCheck(err)

	//logger.Debug("SUCCESS disconnectMongo")
}
