package liveleaderboard

/*
CAUTION - This is a stateful portion of the backend, and should not be used if the backend is to be scaled to
multiple instances. In that case, implement similar functionality in a different service that is common to all the
instances, such as a redis based leaderboard in a separate container.
*/

import (
	"fmt"
	"sort"
	"time"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/sirupsen/logrus"
)

// NewUser is used to add a new user to the in-memory leaderboard
var NewUser chan database.LeaderboardEntry

// UpdateUser is used to update the in-memory leaderboard
var UpdateUser chan int

// FetchRank is used to fetch the rank of a user from the in-memory leaderboard
var FetchRank chan int

// ReturnRank is used to return the fetched rank back to the main routine
var ReturnRank chan int

func sortLeaderboard(leaderboard []database.LeaderboardEntry) {
	sort.Slice(leaderboard, func(i, j int) bool {
		if leaderboard[i].CurrLevel == leaderboard[j].CurrLevel {
			return leaderboard[i].LastAnsTime.Before(leaderboard[j].LastAnsTime)
		}
		return leaderboard[i].CurrLevel > leaderboard[j].CurrLevel
	})
}

// InitLiveLeaderboard initializes the channels and starts the goroutine that maintains the in-memory leaderboard
func InitLiveLeaderboard(db *database.DB) {
	NewUser = make(chan database.LeaderboardEntry)
	UpdateUser = make(chan int)
	FetchRank = make(chan int)
	var leaderboard []database.LeaderboardEntry
	go func() {
		// Populate leaderboard from DB
		err := db.GetLeaderboardData(&leaderboard)
		if err != nil && err.Error() != "sql: no rows in result set" {
			logrus.Info(err)
			logrus.Info("Failed to init in-memory leaderboard")
			return
		}
		sortLeaderboard(leaderboard)

		for {
			select {
			case user := <-NewUser:
				fmt.Println("Adding new user to leaderboard")
				leaderboard = append(leaderboard, user)
				sortLeaderboard(leaderboard)
				for _, v := range leaderboard {
					fmt.Println(v.CurrLevel)
				}

			case userID := <-UpdateUser:
				fmt.Println("Updating leaderboard")
				for i := range leaderboard {
					if leaderboard[i].Uid == userID {
						leaderboard[i].CurrLevel++
						leaderboard[i].LastAnsTime = time.Now()
						break
					}
				}
				sortLeaderboard(leaderboard)
				for _, v := range leaderboard {
					fmt.Println(v.CurrLevel)
				}

			case requestedUID := <-FetchRank:
				fmt.Println("Fetching rank")
				uidFound := false
				for i, v := range leaderboard {
					if v.Uid == requestedUID {
						fmt.Println(i + 1)
						uidFound = true
						break
					}
				}
				if !uidFound {
					fmt.Println("invalid")
				}
			}
		}
	}()
}
