package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func InitDB() error {
	// Parse connection string and disable prepared statement cache for transaction pooler
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	// Disable prepared statement cache (required for Supabase transaction pooler)
	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}

	// Test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		return err
	}

	Conn = conn
	log.Println("Connected to:", version)
	return nil
}

func Close() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}
