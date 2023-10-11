package api

import (
	v1 "shop-api/api/controller/v1"

	"github.com/labstack/echo/v4"
)

func InitRoutes() {
	e := echo.New()

	userController := v1.InitUserController()
	productController := v1.InitProductController()

	e.POST("/users", userController.CreateUser)
	e.POST("/products", productController.CreateProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
