package app

import (
	"shop-api/api"
	"shop-api/db"
	"shop-api/handlers"
	"shop-api/services"
)

type Application struct {
}

// Returns a new instance of the application
func NewApplication() Application {
	return Application{}
}

func (a Application) InitApplication() {
	database := db.InitConnection()
	h := handlers.New(database)

	services.InitServices(h)
	api.InitRoutes()

	db.CloseConnection(database)
}
