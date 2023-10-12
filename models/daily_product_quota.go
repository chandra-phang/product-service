package models

import "time"

type DailyProductQuota struct {
	ID          string
	ProductID   string
	DailyQuota  int
	BookedQuota int
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
