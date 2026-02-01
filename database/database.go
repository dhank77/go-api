package database

import (
	"context"

	"go-api/config"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect(cfg *config.Config) error {
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return err
	}

	DB = conn
	return nil
}

func Close() {
	if DB != nil {
		DB.Close(context.Background())
	}
}
