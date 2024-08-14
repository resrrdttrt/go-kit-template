package db

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

var _ Database = (*database)(nil)

type database struct {
	dbWrite *sqlx.DB
	dbRead  *sqlx.DB
}

// Database provides a database interface
type Database interface {
	HealthCheck(context.Context) error
	NamedExecContext(context.Context, string, interface{}) (sql.Result, error)
	QueryRowxContext(context.Context, string, ...interface{}) *sqlx.Row
	NamedQueryContext(context.Context, string, interface{}) (*sqlx.Rows, error)
	NamedExecWithResponse(context.Context, string, interface{}) (*sqlx.Rows, error)
	GetContext(context.Context, interface{}, string, ...interface{}) error
	BeginTxx(context.Context, *sql.TxOptions) (*sqlx.Tx, error)
	PrepareIn(string, ...interface{}) (string, []interface{}, error)
}


// NewReadWrite create a Database instance with separated read/write client
func NewReadWrite(rdb *sqlx.DB, wdb *sqlx.DB) Database {
	return &database{
		dbWrite: wdb,
		dbRead:  rdb,
	}
}

func (dm database) PrepareIn(s string, i ...interface{}) (string, []interface{}, error) {
	q, args, err := sqlx.In(s, i...)
	if err != nil {
		return q, args, err
	}
	return dm.dbWrite.Rebind(q), args, nil
}

func (dm database) HealthCheck(ctx context.Context) error {
	return dm.dbWrite.PingContext(ctx)
}

func (dm database) NamedExecContext(ctx context.Context, query string, args interface{}) (sql.Result, error) {
	addSpanTags(ctx, query)
	return dm.dbWrite.NamedExecContext(ctx, query, args)
}

func (dm database) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	addSpanTags(ctx, query)
	return dm.dbRead.QueryRowxContext(ctx, query, args...)
}

func (dm database) NamedQueryContext(ctx context.Context, query string, args interface{}) (*sqlx.Rows, error) {
	addSpanTags(ctx, query)
	return dm.dbRead.NamedQueryContext(ctx, query, args)
}

func (dm database) NamedExecWithResponse(ctx context.Context, query string, args interface{}) (*sqlx.Rows, error) {
	addSpanTags(ctx, query)
	return dm.dbWrite.NamedQueryContext(ctx, query, args)
}

func (dm database) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	addSpanTags(ctx, query)
	return dm.dbRead.GetContext(ctx, dest, query, args...)
}

func (dm database) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("span.kind", "client")
		span.SetTag("peer.service", "postgres")
		span.SetTag("dbWrite.type", "sql")
	}
	return dm.dbWrite.BeginTxx(ctx, opts)
}

func addSpanTags(ctx context.Context, query string) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("sql.statement", query)
		span.SetTag("span.kind", "client")
		span.SetTag("peer.service", "postgres")
		span.SetTag("dbWrite.type", "sql")
	}
}

