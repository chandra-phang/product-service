package model

import (
	"product-service/lib"
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID         string
	Name       string
	DailyQuota int
	Status     ProductStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ProductStatus string

const (
	ProductEnabled  ProductStatus = "ENABLED"
	ProductDisabled ProductStatus = "DISABLED"
)

func (Product) Initialize(name string, dailyQuota int) *Product {
	return &Product{
		ID:         lib.GenerateUUID(),
		Name:       name,
		DailyQuota: dailyQuota,
		Status:     ProductEnabled,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type IProductRepository interface {
	CreateProduct(ctx echo.Context, product Product) error
	ListProducts(ctx echo.Context) ([]Product, error)
	GetProduct(ctx echo.Context, productID string) (*Product, error)
	UpdateProduct(ctx echo.Context, product Product) error
	DisableProduct(ctx echo.Context, productID string) error
	EnableProduct(ctx echo.Context, productID string) error
}
