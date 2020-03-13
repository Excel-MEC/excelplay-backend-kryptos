package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func setupDatabase(db *sqlx.DB) error {
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
	);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return errors.Wrap(err, "Could not create schema")
	}

	return nil
}
