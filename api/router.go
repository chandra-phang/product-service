package api

import (
	v1 "shop-api/api/controller/v1"

	"github.com/labstack/echo/v4"
)

func InitRoutes() {
	e := echo.New()

	userController := v1.InitUserController()

	e.POST("/users", userController.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
