package services

import (
	"database/sql"
	"product-service/apperrors"
	"product-service/db"
	v1request "product-service/dto/request/v1"
	"product-service/handlers"
	"product-service/model"
	"product-service/repositories"
	"time"

	"github.com/labstack/echo/v4"
)

type IProductService interface {
	// svc CRUD methods for domain objects
	ListProducts(ctx echo.Context) ([]model.Product, error)
	CreateProduct(ctx echo.Context, dto v1request.CreateProductDTO) error
	GetProduct(ctx echo.Context, productID string) (*model.Product, error)
	UpdateProduct(ctx echo.Context, productID string, dto v1request.UpdateProductDTO) error
	DisableProduct(ctx echo.Context, productID string) error
	EnableProduct(ctx echo.Context, productID string) error
	IncreaseBookedQuota(ctx echo.Context, dto v1request.IncreaseBookedQuotaDTO) error
	DecreaseBookedQuota(ctx echo.Context, dto v1request.DecreaseBookedQuotaDTO) error
}

type productSvc struct {
	dbCon                 *sql.DB
	productRepo           model.IProductRepository
	dailyProductQuotaRepo model.IDailyProductQuotaRepository
}

var productSvcSingleton IProductService

func InitProductService(h handlers.Handler) {
	productSvcSingleton = productSvc{
		dbCon:                 db.GetDB(),
		productRepo:           repositories.NewProductRepositoryInstance(h.DB),
		dailyProductQuotaRepo: repositories.NewDailyProductQuotaRepositoryInstance(h.DB),
	}
}

func GetProductService() IProductService {
	return productSvcSingleton
}

func (svc productSvc) CreateProduct(ctx echo.Context, dto v1request.CreateProductDTO) error {
	product := new(model.Product).Initialize(dto.Name, dto.DailyQuota)
	err := svc.productRepo.CreateProduct(ctx, *product)
	if err != nil {
		return err
	}
	return nil
}

func (svc productSvc) ListProducts(ctx echo.Context) ([]model.Product, error) {
	products, err := svc.productRepo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (svc productSvc) GetProduct(ctx echo.Context, productID string) (*model.Product, error) {
	product, err := svc.productRepo.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (svc productSvc) UpdateProduct(ctx echo.Context, productID string, dto v1request.UpdateProductDTO) error {
	product, err := svc.productRepo.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	product.Name = dto.Name
	product.DailyQuota = dto.DailyQuota
	product.UpdatedAt = time.Now()

	err = svc.productRepo.UpdateProduct(ctx, *product)
	if err != nil {
		return err
	}

	return nil
}

func (svc productSvc) DisableProduct(ctx echo.Context, productID string) error {
	product, err := svc.productRepo.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	if product.Status == model.ProductDisabled {
		return apperrors.ErrProductAlreadyDisabled
	}

	err = svc.productRepo.DisableProduct(ctx, productID)
	if err != nil {
		return err
	}

	return nil
}

func (svc productSvc) EnableProduct(ctx echo.Context, productID string) error {
	product, err := svc.productRepo.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	if product.Status == model.ProductEnabled {
		return apperrors.ErrProductAlreadyEnabled
	}

	err = svc.productRepo.EnableProduct(ctx, productID)
	if err != nil {
		return err
	}

	return nil
}

func (svc productSvc) IncreaseBookedQuota(ctx echo.Context, dto v1request.IncreaseBookedQuotaDTO) error {
	// add DB transaction
	tx, _ := svc.dbCon.Begin()
	defer tx.Rollback()

	for _, productDTO := range dto.Products {
		product, err := svc.productRepo.GetProduct(ctx, productDTO.ProductID)
		if err != nil {
			return err
		}

		dailyProductQuota, err := svc.dailyProductQuotaRepo.GetDailyProductQuota(ctx, product.ID, time.Now())
		if err != nil && err != apperrors.ErrDailyProductQuotaNotFound {
			return err
		}

		if dailyProductQuota == nil {
			dailyProductQuota = new(model.DailyProductQuota).Initialize(product.ID, product.DailyQuota)
			err = svc.dailyProductQuotaRepo.CreateDailyProductQuota(ctx, *dailyProductQuota)
			if err != nil {
				return err
			}
		}

		bookQuantity := productDTO.Quantity
		if dailyProductQuota.BookedQuota+bookQuantity >= dailyProductQuota.DailyQuota {
			return apperrors.ErrProductBookedQuotaReachLimit
		}

		err = svc.dailyProductQuotaRepo.IncreaseDailyProductQuota(ctx, dailyProductQuota.ID, bookQuantity)
		if err != nil {
			return err
		}
	}

	// commit the transaction
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (svc productSvc) DecreaseBookedQuota(ctx echo.Context, dto v1request.DecreaseBookedQuotaDTO) error {
	// add DB transaction
	tx, _ := svc.dbCon.Begin()
	defer tx.Rollback()

	for _, productDTO := range dto.Products {
		product, err := svc.productRepo.GetProduct(ctx, productDTO.ProductID)
		if err != nil {
			return err
		}

		dailyProductQuota, err := svc.dailyProductQuotaRepo.GetDailyProductQuota(ctx, product.ID, time.Now())
		if err != nil {
			return err
		}

		quantity := productDTO.Quantity
		if dailyProductQuota.BookedQuota-quantity <= 0 {
			return apperrors.ErrProductBookedQuotaCannotDecrease
		}

		err = svc.dailyProductQuotaRepo.DecreaseDailyProductQuota(ctx, dailyProductQuota.ID, quantity)
		if err != nil {
			return err
		}
	}

	// commit the transaction
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
