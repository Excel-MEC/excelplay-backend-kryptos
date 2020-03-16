package database

import (
	"fmt"

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
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Dbport,
		config.User,
		config.Password,
		config.Dbname,
	)

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to the db")
	}
	defer db.Close()
	// This step is needed because db.Open() simply validates the arguments, it does not open an actual connection to the db.
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Could not ping the db to establish conection")
	}

	// execute a query on the server
	_, err = db.Exec(schema)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create schema")
	}

	return &DB{db}, nil
}
