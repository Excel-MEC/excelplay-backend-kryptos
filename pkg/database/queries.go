package database

import "database/sql"

// GetCurrLevel retuns the current level for the user with the given uuid.
func (db *DB) GetCurrLevel(uuid string, currLev *int) error {
	return db.Get(currLev, "select curr_level from kuser where id = $1", uuid)
}

// CreateNewUser creates a new user, who starts at level 1.
func (db *DB) CreateNewUser(uuid string, name string) (sql.Result, error) {
	return db.Exec("insert into kuser values($1,$2,$3)", uuid, name, 1)
}
