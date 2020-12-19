package database

import (
	"time"
)

// User holds the details of a particular user
type User struct {
	Name      string `json:"name" db:"name"`
	CurrLevel int    `json:"curr_level" db:"curr_level"`
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

// LeaderboardEntry holds an entry in the in-memory leaderboard
type LeaderboardEntry struct {
	uid         int
	currLevel   int
	lastAnsTime time.Time
}
