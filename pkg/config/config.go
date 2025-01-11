package config

import "github.com/jackc/pgx/v4"

// AppConfig holds app configurations
type AppConfig struct {
	InProduction bool
	DB           *pgx.Conn
}
