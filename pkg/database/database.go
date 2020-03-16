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
func NewDB(config *env.Config) (*DB, error) {
	schema := `create table if not exists levels (
		number int not null primary key,
		question varchar(1000),
		image_level boolean not null,
		level_file varchar(1000),
		answer varchar(100)
	);

	create table if not exists hints (
		id serial primary key,
		number int references levels(number) not null,
		content varchar(2000)
	);
	
	create extension if not exists "uuid-ossp";

	create table if not exists kuser (
		id uuid primary key default uuid_generate_v1(),
		name varchar(100) not null,
		curr_level int references levels(number) not null,
		last_anstime timestamp
	);
	
	create table if not exists answer_logs (
		id uuid references kuser(id),
		name varchar(100),
		attempt varchar(100),
		time timestamp		
	)
	`
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
