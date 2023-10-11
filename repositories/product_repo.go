package repositories

import (
	"database/sql"
	"shop-api/models"

	"github.com/labstack/echo/v4"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepositoryInstance(db *sql.DB) models.IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r ProductRepository) CreateProduct(ctx echo.Context, product models.Product) error {
	sqlStatement := `
		INSERT INTO products
			(id, name, daily_quota, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?)
	`

	params := []interface{}{
		product.ID,
		product.Name,
		product.DailyQuota,
		product.CreatedAt,
		product.UpdatedAt,
	}

	_, err := r.db.Exec(sqlStatement, params...)
	if err != nil {
		return err
	}

	return nil
}
