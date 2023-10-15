package api

import (
	"log"
	v1 "product-service/api/controllers/v1"

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
	v1Api.PUT("/products/:id/enable", productController.EnableProduct)
	v1Api.PUT("/products/:id/increase-booked-quota", productController.IncreaseBookedQuota)
	v1Api.PUT("/products/:id/decrease-booked-quota", productController.DecreaseBookedQuota)

	log.Println("Server is running at 8080 port.")
	e.Logger.Fatal(e.Start(":8080"))
}
