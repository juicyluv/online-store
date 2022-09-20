package database

import (
	"context"
)

type Rows interface {
	Scan(dest ...any) error
	Next() bool
	Err() error
	Close()
}

type Database interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (int64, error)
	Close(ctx context.Context) error
}
