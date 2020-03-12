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
		level_file varchar(1000)
	);

	create table if not exists hints (
		id serial primary key,
		number int references levels(number) not null,
		content varchar(1000)
	);
	
	create table if not exists leaderboard (
		uid int primary key,
		name varchar(100) not null,
		curr_level int references levels(number) not null,
		rank int not null
	);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return errors.Wrap(err, "Could not create schema")
	}

	return nil
}
