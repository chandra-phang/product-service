package model

import (
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

type IProductRepository interface {
	CreateProduct(ctx echo.Context, product Product) error
	ListProducts(ctx echo.Context) ([]Product, error)
	GetProduct(ctx echo.Context, productID string) (*Product, error)
	UpdateProduct(ctx echo.Context, product Product) error
	DisableProduct(ctx echo.Context, productID string) error
	EnableProduct(ctx echo.Context, productID string) error
}
