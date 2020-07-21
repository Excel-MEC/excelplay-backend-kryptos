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

// QResponse holds the response for a question request
type QResponse struct {
	// Values that can be nil or a non-nullable value,
	// such as a string are given the empty interface type
	Number     int         `json:"number" db:"number"`
	Question   interface{} `json:"question" db:"question"`
	ImageLevel bool        `json:"image_level" db:"image_level"`
	LevelFile  interface{} `json:"level_file" db:"level_file"`
	Hints      []string    `json:"hints"`
}

// GetQuestion gets details of a certain level
func (db *DB) GetQuestion(currLev int, res *QResponse) error {
	return db.Get(res, "select number, question, image_level, level_file from levels where number = $1", currLev)
}

// GetHints gets the hints released for a question
func (db *DB) GetHints(currLev int, hints *[]string) error {
	return db.Select(hints, "select content from hints where number = $1", currLev)
}
