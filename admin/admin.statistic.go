package admin

import (
	"context"
	"time"
)

type Statistic struct {
	Day   time.Time `json:"day"`
	Count int       `json:"count"`
}

type StatisticRepository interface {
	GetTotalUsers(ctx context.Context) (int, error)
	GetTotalGames(ctx context.Context) (int, error)
	GetTotalEnterprises(ctx context.Context) (int, error)
	GetTotalEndUser(ctx context.Context) (int, error)
	GetTotalActiveEndUsers(ctx context.Context) (int, error)
	GetTotalActiveEnterprises(ctx context.Context) (int, error)
	GetTotalNewEnterprisesInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error)
	GetTotalNewEndUsersInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error)
}
