package database

import (
	"fmt"
	"strconv"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// DB wraps sqlx.DB to add custom query methods
type DB struct {
	*sqlx.DB
}

// NewDB setups and returns a new database connection instance
func NewDB(config *env.DBConfig) (*DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Dbport,
		config.User,
		config.Password,
		config.Dbname,
		config.SSLMode,
	)

	// Run Migrations
	err := Migrate("/excelplay-backend-kryptos/pkg/database/migrations", &DBParams{config.Host, strconv.Itoa(config.Dbport), config.User, config.Password, config.Dbname, config.SSLMode}, "postgres")
	if err != nil {
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

	return &DB{db}, nil
}
