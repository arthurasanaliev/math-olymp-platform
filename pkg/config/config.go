package config

import "github.com/arthurasanaliev/math-olymp-platform/pkg/db"

// AppConfig holds app configurations
type AppConfig struct {
	InProduction bool
	DB           *db.DB
}
