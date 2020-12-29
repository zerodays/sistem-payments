package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/zerodays/sistem-payments/internal/config"
	"github.com/zerodays/sistem-payments/internal/logger"
	"path/filepath"
)

var DB *sqlx.DB

// Init initializes the database.
func Init() {
	// Create database URL.
	databaseUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host(),
		config.Database.Port(),
		config.Database.User(),
		config.Database.Password(),
		config.Database.DbName(),
		config.Database.SslMode())

	// Open DB connection.
	var err error
	DB, err = sqlx.Open("pgx", databaseUrl)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}

	// Ping DB to check the connection.
	err = DB.Ping()
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}

	// Migrate database to the latest version.
	migrateDatabase()
}

// migrate migrates to the latest version.
func migrateDatabase() {
	// Get migrations.
	assets, err := config.AssetDir("migrations")
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}

	// Create bindata resources.
	resources := bindata.Resource(assets,
		func(name string) ([]byte, error) {
			return config.Asset(filepath.Join("migrations", name))
		})

	// Create migration source instance.
	data, err := bindata.WithInstance(resources)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}

	// Create database driver.
	driver, err := postgres.WithInstance(DB.DB, &postgres.Config{})
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}

	// Create new migration object.
	migration, err := migrate.NewWithInstance("go-bindata", data, "postgres", driver)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}

	// Migrate database to the newest migration.
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Log.Fatal().Err(err).Send()
	}
}

// Close closes the database.
func Close() {
	err := DB.Close()
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}
}
