package model

import (
	"time"

	"github.com/labstack/echo/v4"
)

type DailyProductQuota struct {
	ID          string
	ProductID   string
	DailyQuota  int
	BookedQuota int
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type IDailyProductQuotaRepository interface {
	CreateDailyProductQuota(ctx echo.Context, dailyProductQuota DailyProductQuota) error
	GetDailyProductQuota(ctx echo.Context, productID string, date time.Time) (*DailyProductQuota, error)
	IncreaseDailyProductQuota(ctx echo.Context, dailyProductQuotaID string, quantity int) error
	DecreaseDailyProductQuota(ctx echo.Context, dailyProductQuotaID string, quantity int) error
}
