package liveleaderboard

/*
CAUTION - This is a stateful portion of the backend, and should not be used if the backend is to be scaled to
multiple instances. In that case, implement similar functionality in a different service that is common to all the
instances, such as a redis based leaderboard.
*/

import (
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
)

// NewUser is used to add a new user to the in-memory leaderboard
var NewUser chan database.LeaderboardEntry

// UpdateUser is used to update the in-memory leaderboard
var UpdateUser chan database.LeaderboardEntry

// FetchRank is used to fetch the rank of a user from the in-memory leaderboard
var FetchRank chan int

// InitLiveLeaderboard initializes the channels and starts the goroutine that maintains the in-memory leaderboard
func InitLiveLeaderboard() {

}
