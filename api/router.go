package api

import (
	v1 "shop-api/api/controller/v1"

	"github.com/labstack/echo/v4"
)

func InitRoutes() {
	e := echo.New()

	productController := v1.InitProductController()

	v1Api := e.Group("v1")
	v1Api.GET("/products", productController.ListProducts)
	v1Api.POST("/products", productController.CreateProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
