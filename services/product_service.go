package services

import (
	"product-service/apperrors"
	v1request "product-service/dto/request/v1"
	"product-service/handlers"
	"product-service/lib"
	"product-service/models"
	"product-service/repositories"
	"time"

	"github.com/labstack/echo/v4"
)

type IProductService interface {
	// svc CRUD methods for domain objects
	ListProducts(ctx echo.Context) ([]models.Product, error)
	CreateProduct(ctx echo.Context, dto v1request.CreateProductDTO) error
	GetProduct(ctx echo.Context, productID string) (*models.Product, error)
	UpdateProduct(ctx echo.Context, productID string, dto v1request.UpdateProductDTO) error
	DisableProduct(ctx echo.Context, productID string) error
	EnableProduct(ctx echo.Context, productID string) error
	IncreaseBookedQuota(ctx echo.Context, productiD string) error
	DecreaseBookedQuota(ctx echo.Context, productiD string) error
}

type productSvc struct {
	ProductRepo           models.IProductRepository
	DailyProductQuotaRepo models.IDailyProductQuotaRepository
}

var productSvcSingleton IProductService

func InitProductService(h handlers.Handler) {
	productSvcSingleton = productSvc{
		ProductRepo:           repositories.NewProductRepositoryInstance(h.DB),
		DailyProductQuotaRepo: repositories.NewDailyProductQuotaRepositoryInstance(h.DB),
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

func (svc productSvc) DisableProduct(ctx echo.Context, productID string) error {
	product, err := svc.ProductRepo.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	if product.Status == models.ProductDisabled {
		return apperrors.ErrProductAlreadyDisabled
	}

	err = svc.ProductRepo.DisableProduct(ctx, productID)
	if err != nil {
		return err
	}

	return nil
}

func (svc productSvc) EnableProduct(ctx echo.Context, productID string) error {
	product, err := svc.ProductRepo.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	if product.Status == models.ProductEnabled {
		return apperrors.ErrProductAlreadyEnabled
	}

	err = svc.ProductRepo.EnableProduct(ctx, productID)
	if err != nil {
		return err
	}

	return nil
}

func (svc productSvc) IncreaseBookedQuota(ctx echo.Context, productID string) error {
	product, err := svc.ProductRepo.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	dailyProductQuota, err := svc.DailyProductQuotaRepo.GetDailyProductQuota(ctx, product.ID, time.Now())
	if err != nil && err != apperrors.ErrDailyProductQuotaNotFound {
		return err
	}

	if dailyProductQuota == nil {
		dailyProductQuota = &models.DailyProductQuota{
			ID:          lib.GenerateUUID(),
			ProductID:   product.ID,
			DailyQuota:  product.DailyQuota,
			BookedQuota: 0,
			Date:        time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		err = svc.DailyProductQuotaRepo.CreateDailyProductQuota(ctx, *dailyProductQuota)
		if err != nil {
			return err
		}
	}

	if dailyProductQuota.BookedQuota >= dailyProductQuota.DailyQuota {
		return apperrors.ErrProductBookedQuotaReachLimit
	}

	err = svc.DailyProductQuotaRepo.IncreaseDailyProductQuota(ctx, dailyProductQuota.ID)
	if err != nil {
		return err
	}

	return nil
}

func (svc productSvc) DecreaseBookedQuota(ctx echo.Context, productID string) error {
	product, err := svc.ProductRepo.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	dailyProductQuota, err := svc.DailyProductQuotaRepo.GetDailyProductQuota(ctx, product.ID, time.Now())
	if err != nil {
		return err
	}

	if dailyProductQuota.BookedQuota <= 0 {
		return apperrors.ErrProductBookedQuotaCannotDecrease
	}

	err = svc.DailyProductQuotaRepo.DecreaseDailyProductQuota(ctx, dailyProductQuota.ID)
	if err != nil {
		return err
	}

	return nil
}
