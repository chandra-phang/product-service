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
	v1Api.GET("/products/:id", productController.GetProduct)
	v1Api.PUT("/products/:id", productController.UpdateProduct)
	v1Api.PUT("/products/:id/disable", productController.DisableProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
