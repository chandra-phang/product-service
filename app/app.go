package app

import (
	"product-service/api"
	"product-service/db"
	"product-service/handlers"
	"product-service/services"
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
