package database

// GetCurrLevel retruns the current level for the user with the given uuid
func (db *DB) GetCurrLevel(uuid string) error {
	var currLev int
	return db.Get(&currLev, "select curr_level from kuser where id = $1", uuid)
}
