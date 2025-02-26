package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %w", err)
	}

	migrationPath := filepath.Join("../../internal", "db", "migrations")
	if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
		return fmt.Errorf("migrations directory does not exist: %s", migrationPath)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../internal/db/migrations",
		"sqlite3", driver)
	if err != nil {
		return fmt.Errorf("could not create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	return nil
}
