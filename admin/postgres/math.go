package postgres

import (
	"context"

	"go-kit-template/admin"
	"go-kit-template/pkg/db"
	log "go-kit-template/pkg/logger"
)

var _ admin.MathRepository = (*mathRepository)(nil)	

type mathRepository struct {
	db db.Database
	l  log.Logger
}

func NewMathRepository(db db.Database, l log.Logger) admin.MathRepository {
	return &mathRepository{
		db: db,
		l:  l,
	}
}

func (r *mathRepository) Add(ctx context.Context, numA, numB float32) (sum float32, err error) {
	return numA + numB, nil
}