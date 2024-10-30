package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*bun.DB
}

// Dial creates new database connection to postgres
func Dial(dsn string) (*DB, error) {
	if dsn == "" {
		return nil, errors.New("no postgres DSN provided")
	}

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(1 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping failed: %w", err)
	}

	bunDB := bun.NewDB(db, pgdialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return &DB{bunDB}, nil
}
