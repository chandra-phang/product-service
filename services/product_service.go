package services

import (
	v1request "shop-api/dto/request/v1"
	"shop-api/handlers"
	"shop-api/lib"
	"shop-api/models"
	"shop-api/repositories"
	"time"

	"github.com/labstack/echo/v4"
)

type IProductService interface {
	// svc CRUD methods for domain objects
	CreateProduct(ctx echo.Context, dto v1request.CreateProductDTO) error
}

type productSvc struct {
	ProductRepo models.IProductRepository
}

var productSvcSingleton IProductService

func InitProductService(h handlers.Handler) {
	productSvcSingleton = productSvc{
		ProductRepo: repositories.NewProductRepositoryInstance(h.DB),
	}
}

func GetProductService() IProductService {
	return productSvcSingleton
}

func (svc productSvc) CreateProduct(ctx echo.Context, dto v1request.CreateProductDTO) error {
	product := models.Product{
		ID:         lib.GenerateUUID(),
		Name:       dto.Name,
		DailyQuota: dto.DailyQuota,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := svc.ProductRepo.CreateProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
