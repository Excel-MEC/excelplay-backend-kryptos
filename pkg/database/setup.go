package database

import (
	"fmt"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// DB wraps sqlx.DB to add custom query methods
type DB struct {
	*sqlx.DB
}

// NewDB setups and returns a new database connection instance
func NewDB(config *env.DBConfig) (*DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Dbport,
		config.User,
		config.Password,
		config.Dbname,
	)

	// Read migrations
	m, err := migrate.New("file:///excelplay-backend-kryptos/pkg/database/migrations", "postgres://admin:password@db:5432/db?sslmode=disable")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read migrations or connect to db")
	}
	// The error happens when running m.Up()
	// Run migrations
	if err := m.Up(); err != nil {
		return nil, errors.Wrap(err, "Failed to run migrations")
	}

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil || db == nil {
		return nil, errors.Wrap(err, "Could not connect to the db")
	}

	// This step is needed because db.Open() simply validates the arguments, it does not open an actual connection to the db.
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Could not ping the db to establish conection")
	}

	// execute a query on the server
	// _, err = db.Exec(schema)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Could not create schema")
	// }

	return &DB{db}, nil
}
