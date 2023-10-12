package repositories

import (
	"database/sql"
	"log"
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

func (r ProductRepository) ListProducts(ctx echo.Context) ([]models.Product, error) {
	sqlStatement := `
		SELECT
			id,
			name,
			daily_quota,
			created_at,
			updated_at
		FROM products
		ORDER BY updated_at DESC
	`

	results, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var products = make([]models.Product, 0)
	for results.Next() {
		var product models.Product
		err = results.Scan(&product.ID, &product.Name, &product.DailyQuota, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}
