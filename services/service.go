package services

import "shop-api/handlers"

func InitServices(h handlers.Handler) {
	InitUserService(h)
	InitProductService(h)
}
