package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Postres struct {
	db *sqlx.DB
}

func NewPostgres(cfg *config.Config, lc fx.Lifecycle) *Postres {
	logger.Info("Initializing to PostgreSQL client...")
	source := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.Creds.DB.User, cfg.Creds.DB.Password, cfg.Creds.DB.Name)
	db, err := sqlx.Connect("postgres", source)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)

	pg := &Postres{db: db}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return pg.RunMigrations()
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return pg
}

func (p *Postres) RunMigrations() error {
	driver, err := migratePostgres.WithInstance(p.db.DB, &migratePostgres.Config{})
	if err != nil {
		return fmt.Errorf("create migrate driver: %w", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	migrationPath := filepath.Join(wd, "migrations")

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("init migrations: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	logger.Info("Database migrations applied successfully")
	return nil
}
