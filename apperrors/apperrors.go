package apperrors

import "errors"

var (
	ErrProductNotFound                = errors.New("product not found")
	ErrProductAlreadyDisabled         = errors.New("product is already disabled")
	ErrProductAlreadyEnabled          = errors.New("product is already enabled")
	ErrDailyProductQuotaNotFound      = errors.New("daily_product_quota not found")
	ErrProductBookedQuotaLimitReached = errors.New("product booked_quota limit reached")
)
