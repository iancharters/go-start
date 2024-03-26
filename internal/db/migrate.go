package db

import (
	"context"
	"embed"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // required for applying migrations to postgres database.
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*
var migrationContents embed.FS

func migrationsSource() (source.Driver, error) {
	return iofs.New(migrationContents, "migrations")
}

// Migrate migrates the db up to the latest version.
func Migrate(ctx context.Context, dsn string) error {
	sourceInstance, err := migrationsSource()
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("embed", sourceInstance, dsn)
	if err != nil {
		return err
	}
	defer migrations.Close()

	if err := migrations.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
