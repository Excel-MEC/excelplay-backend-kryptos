package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	// if any error occurs during startup, log the error and exit with status 1
	if err := startup(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
func startup() error {
	//setup logger
	logger := logrus.New()
	logger.Out = os.Stdout

	//setup router
	router := mux.NewRouter()

	//setup the database
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, DBPORT, USER, PASSWORD, DBNAME)
	// db, err := sqlx.Connect("sqlite3", ":memory:")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return errors.Wrap(err, "Could not setup the db")
	}
	defer db.Close()
	// This step is needed because db.Open() simply validates the arguments, it does not open an actual connection to the db.
	err = db.Ping()
	if err != nil {
		return err
	}

	server := &http.Server{
		Handler:      newServer(router, db, logger),
		Addr:         PORT,
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	//start server
	logger.Info("Server starting on port " + PORT)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return errors.Wrap(err, "Could not start server on port "+PORT)
	}
	return nil
}
