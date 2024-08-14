package postgres

import (
	"context"
	"time"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/db"
	"github.com/resrrdttrt/VOU/pkg/errors"
	log "github.com/resrrdttrt/VOU/pkg/logger"
)

var _ admin.StatisticRepository = (*statisticRepository)(nil)

type statisticRepository struct {
	db db.Database
	l  log.Logger
}

func NewStatisticRepository(db db.Database, l log.Logger) admin.StatisticRepository {
	return &statisticRepository{
		db: db,
		l:  l,
	}
}

func (r *statisticRepository) GetTotalUsers(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users`
	params := map[string]interface{}{}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return 0, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, errors.Wrap(ErrSelectDb, err)
		}
		return count, nil
	} else {
		return 0, nil
	}
}

func (r *statisticRepository) GetTotalGames(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM games`
	params := map[string]interface{}{}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return 0, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, errors.Wrap(ErrSelectDb, err)
		}
		return count, nil
	} else {
		return 0, nil
	}
}

func (r *statisticRepository) GetTotalEnterprises(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users where role = 'enterprise'`
	params := map[string]interface{}{}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return 0, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, errors.Wrap(ErrSelectDb, err)
		}
		return count, nil
	} else {
		return 0, nil
	}
}

func (r *statisticRepository) GetTotalEndUser(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users where role = 'end_user'`
	params := map[string]interface{}{}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return 0, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, errors.Wrap(ErrSelectDb, err)
		}
		return count, nil
	} else {
		return 0, nil
	}
}

func (r *statisticRepository) GetTotalActiveEndUsers(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users where role = 'end_user' and status = 'active'`
	params := map[string]interface{}{}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return 0, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, errors.Wrap(ErrSelectDb, err)
		}
		return count, nil
	} else {
		return 0, nil
	}
}

func (r *statisticRepository) GetTotalActiveEnterprises(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users where role = 'enterprise' and status = 'active'`
	params := map[string]interface{}{}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return 0, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, errors.Wrap(ErrSelectDb, err)
		}
		return count, nil
	} else {
		return 0, nil
	}
}

func (r *statisticRepository) GetTotalNewEnterprisesInTime(ctx context.Context, start time.Time, end time.Time) ([]admin.Statistic, error) {
	query := `SELECT DATE(created_at) AS day, COUNT(id) AS count FROM users where role = 'enterprise' and created_at >= :start and created_at <= :end GROUP BY DATE(created_at) ORDER BY DATE(created_at)`
	params := map[string]interface{}{
		"start": start,
		"end":   end,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return []admin.Statistic{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var statisticResults []admin.Statistic

	for rows.Next() {
		var statistic admin.Statistic
		if err := rows.StructScan(&statistic); err != nil {
			return []admin.Statistic{}, errors.Wrap(ErrSelectDb, err)
		}
		statistic.Day = statistic.Day.Truncate(24 * time.Hour)
		statisticResults = append(statisticResults, statistic)
	}

	if err := rows.Err(); err != nil {
		return []admin.Statistic{}, errors.Wrap(ErrSelectDb, err)
	}
	return statisticResults, nil
}

func (r *statisticRepository) GetTotalNewEndUsersInTime(ctx context.Context, start time.Time, end time.Time) ([]admin.Statistic, error) {
	query := `SELECT DATE(created_at) AS day, COUNT(id) AS count FROM users where role = 'end_user' and created_at >= :start and created_at <= :end GROUP BY DATE(created_at) ORDER BY DATE(created_at)`
	params := map[string]interface{}{
		"start": start,
		"end":   end,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return []admin.Statistic{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var statisticResults []admin.Statistic

	for rows.Next() {
		var statistic admin.Statistic
		if err := rows.StructScan(&statistic); err != nil {
			return []admin.Statistic{}, errors.Wrap(ErrSelectDb, err)
		}
		statistic.Day = statistic.Day.Truncate(24 * time.Hour)
		statisticResults = append(statisticResults, statistic)
	}

	if err := rows.Err(); err != nil {
		return []admin.Statistic{}, errors.Wrap(ErrSelectDb, err)
	}
	return statisticResults, nil
}
