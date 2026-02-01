package database

import (
	"context"
	"fmt"

	"go-api/config"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
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
