package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/juicyluv/online-store/internal/infrastructure/database"
)

type postgresql struct {
	conn *pgx.Conn
}

// 	NewPostgresConnection parses connString and tries to connect to database.
// 	Under the hood it uses ping to database to make sure that connection was successful.
func NewPostgresConnection(ctx context.Context, connString string) (database.Database, error) {
	cfg, err := pgx.ParseConfig(connString)

	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(ctx, cfg)

	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)

	if err != nil {
		return nil, err
	}

	return &postgresql{conn: conn}, nil
}

// 	Query sends the SQL query to the database with given arguments.
// 	It will return rows data as the result or an error if something went wrong.
func (p *postgresql) Query(ctx context.Context, sql string, args ...any) (database.Rows, error) {
	return p.conn.Query(ctx, sql, args)
}

// 	Exec is used to execute some query on database data. It will return count of the affected rows or an error on failure.
// 	Use the Exec for INSERT, UPDATE or DELETE queries that will affect the database data.
func (p *postgresql) Exec(ctx context.Context, sql string, args ...any) (int64, error) {
	result, err := p.conn.Exec(ctx, sql, args)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Close closes a connection. It is safe to call Close on an already closed connection.
func (p *postgresql) Close(ctx context.Context) error {
	return p.conn.Close(ctx)
}
