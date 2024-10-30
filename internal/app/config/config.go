package config

import "os"

// Config is a simple config
type Config struct {
	HTTPAddr       string
	DSN            string
	MigrationsPath string
}

// Read reads config from environment.
func Read() Config {
	var config Config
	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HTTPAddr = httpAddr
	}
	dsn, exists := os.LookupEnv("DSN")
	if exists {
		config.DSN = dsn
	}
	migrationsPath, exists := os.LookupEnv("MIGRATIONS_PATH")
	if exists {
		config.MigrationsPath = migrationsPath
	}
	return config
}
