package utils

import (
	"context"
	"github.com/jackc/pgx/v5"
)

var Pgconn *pgx.Conn

func init() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {

	}
	defer conn.Close(ctx)
}
