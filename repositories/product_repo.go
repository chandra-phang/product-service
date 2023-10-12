package repositories

import (
	"database/sql"
	"errors"
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
			(id, name, daily_quota, status, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?)
	`

	params := []interface{}{
		product.ID,
		product.Name,
		product.DailyQuota,
		product.Status,
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
			status,
			created_at,
			updated_at
		FROM products
		WHERE status = ?
		ORDER BY updated_at DESC
	`

	results, err := r.db.Query(sqlStatement, models.ProductEnabled)
	if err != nil {
		return nil, err
	}

	var products = make([]models.Product, 0)
	for results.Next() {
		var product models.Product
		err = results.Scan(&product.ID, &product.Name, &product.DailyQuota, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (r ProductRepository) GetProduct(ctx echo.Context, productID string) (*models.Product, error) {
	sqlStatement := `
		SELECT
			id,
			name,
			daily_quota,
			status,
			created_at,
			updated_at
		FROM products
		WHERE id = ?
	`

	results, err := r.db.Query(sqlStatement, productID)
	if err != nil {
		return nil, err
	}

	var product models.Product
	for results.Next() {
		err = results.Scan(&product.ID, &product.Name, &product.DailyQuota, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}
	}

	if product.ID == "" {
		return nil, errors.New("product not found")
	}

	return &product, nil
}

func (r ProductRepository) UpdateProduct(ctx echo.Context, product models.Product) error {
	sqlStatement := `
		UPDATE products
		SET
			name = ?,
			daily_quota = ?,
			updated_at = ?
		WHERE id = ?
	`

	params := []interface{}{
		product.Name,
		product.DailyQuota,
		product.UpdatedAt,
		product.ID,
	}

	_, err := r.db.Exec(sqlStatement, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepository) DisableProduct(ctx echo.Context, productID string) error {
	sqlStatement := `
		UPDATE products
		SET
			status = ?
		WHERE id = ?
	`

	params := []interface{}{
		models.ProductDisabled,
		productID,
	}

	_, err := r.db.Exec(sqlStatement, params...)
	if err != nil {
		return err
	}

	return nil
}
