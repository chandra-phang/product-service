package repositories

import (
	"database/sql"
	"log"
	"product-service/apperrors"
	"product-service/model"

	"github.com/labstack/echo/v4"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepositoryInstance(db *sql.DB) model.IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r ProductRepository) CreateProduct(ctx echo.Context, product model.Product) error {
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

func (r ProductRepository) ListProducts(ctx echo.Context) ([]model.Product, error) {
	sqlStatement := `
		SELECT
			p.id,
			p.name,
			p.daily_quota,
			p.status,
			p.created_at,
			p.updated_at
		FROM products p
		WHERE p.status = ?
		ORDER BY p.updated_at DESC
	`

	results, err := r.db.Query(sqlStatement, model.ProductEnabled)
	if err != nil {
		return nil, err
	}

	var products = make([]model.Product, 0)
	for results.Next() {
		var product model.Product
		err = results.Scan(&product.ID, &product.Name, &product.DailyQuota, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (r ProductRepository) GetProduct(ctx echo.Context, productID string) (*model.Product, error) {
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

	var product model.Product
	for results.Next() {
		err = results.Scan(&product.ID, &product.Name, &product.DailyQuota, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}
	}

	if product.ID == "" {
		return nil, apperrors.ErrProductNotFound
	}

	return &product, nil
}

func (r ProductRepository) UpdateProduct(ctx echo.Context, product model.Product) error {
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
		model.ProductDisabled,
		productID,
	}

	_, err := r.db.Exec(sqlStatement, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepository) EnableProduct(ctx echo.Context, productID string) error {
	sqlStatement := `
		UPDATE products
		SET
			status = ?
		WHERE id = ?
	`

	params := []interface{}{
		model.ProductEnabled,
		productID,
	}

	_, err := r.db.Exec(sqlStatement, params...)
	if err != nil {
		return err
	}

	return nil
}
