package models

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	DailyQuota int       `json:"dailyQuota"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type IProductRepository interface {
	CreateProduct(ctx echo.Context, product Product) error
	ListProducts(ctx echo.Context) ([]Product, error)
}
