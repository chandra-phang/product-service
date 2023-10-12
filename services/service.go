package services

import "product-service/handlers"

func InitServices(h handlers.Handler) {
	InitProductService(h)
}
