package repositories

import (
	"database/sql"
	"log"
	"product-service/apperrors"
	"product-service/lib"
	"product-service/model"
	"time"

	"github.com/labstack/echo/v4"
)

type DailyProductQuotaRepository struct {
	db *sql.DB
}

func NewDailyProductQuotaRepositoryInstance(db *sql.DB) model.IDailyProductQuotaRepository {
	return &DailyProductQuotaRepository{
		db: db,
	}
}

func (r DailyProductQuotaRepository) CreateDailyProductQuota(ctx echo.Context, dailyProductQuota model.DailyProductQuota) error {
	sqlStatement := `
		INSERT INTO daily_product_quotas
			(id, product_id, daily_quota, booked_quota, date, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?)
	`

	params := []interface{}{
		dailyProductQuota.ID,
		dailyProductQuota.ProductID,
		dailyProductQuota.DailyQuota,
		dailyProductQuota.BookedQuota,
		dailyProductQuota.Date,
		dailyProductQuota.CreatedAt,
		dailyProductQuota.UpdatedAt,
	}

	_, err := r.db.Exec(sqlStatement, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r DailyProductQuotaRepository) GetDailyProductQuota(ctx echo.Context, productID string, date time.Time) (*model.DailyProductQuota, error) {
	sqlStatement := `
		SELECT
			id,
			product_id,
			daily_quota,
			booked_quota,
			date,
			created_at,
			updated_at
		FROM daily_product_quotas
		WHERE product_id = ? AND date = ?
	`
	dateString := lib.ConvertToDate(date)
	params := []interface{}{productID, dateString}

	results, err := r.db.Query(sqlStatement, params...)
	if err != nil {
		return nil, err
	}

	var dailyProductQuota model.DailyProductQuota
	for results.Next() {
		err = results.Scan(&dailyProductQuota.ID, &dailyProductQuota.ProductID, &dailyProductQuota.DailyQuota, &dailyProductQuota.BookedQuota, &dailyProductQuota.Date, &dailyProductQuota.CreatedAt, &dailyProductQuota.UpdatedAt)
		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}
	}

	if dailyProductQuota.ID == "" {
		return nil, apperrors.ErrDailyProductQuotaNotFound
	}

	return &dailyProductQuota, nil
}

func (r DailyProductQuotaRepository) IncreaseDailyProductQuota(ctx echo.Context, dailyProductQuotaID string) error {
	sqlStatement := `
		UPDATE daily_product_quotas
		SET
			booked_quota = booked_quota + 1
		WHERE id = ?
	`

	_, err := r.db.Exec(sqlStatement, dailyProductQuotaID)
	if err != nil {
		return err
	}

	return nil
}

func (r DailyProductQuotaRepository) DecreaseDailyProductQuota(ctx echo.Context, dailyProductQuotaID string) error {
	sqlStatement := `
		UPDATE daily_product_quotas
		SET
			booked_quota = booked_quota - 1
		WHERE id = ?
	`

	_, err := r.db.Exec(sqlStatement, dailyProductQuotaID)
	if err != nil {
		return err
	}

	return nil
}
