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
	ListProducts(ctx echo.Context) ([]models.Product, error)
	CreateProduct(ctx echo.Context, dto v1request.CreateProductDTO) error
	GetProduct(ctx echo.Context, productID string) (*models.Product, error)
	UpdateProduct(ctx echo.Context, productID string, dto v1request.UpdateProductDTO) error
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
		Status:     models.ProductEnabled,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := svc.ProductRepo.CreateProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (svc productSvc) ListProducts(ctx echo.Context) ([]models.Product, error) {
	products, err := svc.ProductRepo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (svc productSvc) GetProduct(ctx echo.Context, productID string) (*models.Product, error) {
	product, err := svc.ProductRepo.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc productSvc) UpdateProduct(ctx echo.Context, productID string, dto v1request.UpdateProductDTO) error {
	_, err := svc.ProductRepo.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	product := models.Product{
		ID:         productID,
		Name:       dto.Name,
		DailyQuota: dto.DailyQuota,
		UpdatedAt:  time.Now(),
	}

	err = svc.ProductRepo.UpdateProduct(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
