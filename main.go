package main

import (
	config "github.com/kanugi/rest-api-assignment2/app"

	"github.com/kanugi/rest-api-assignment2/controllers"

	"github.com/kanugi/rest-api-assignment2/routers"
)

var PORT = ":8080"

func main() {
	db := config.StartDB()
	controller := controllers.New(db)

	routers.RouteServer(controller).Run(PORT)
}
