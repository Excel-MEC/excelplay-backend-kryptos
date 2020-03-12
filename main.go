package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logger.SetFormatter(formatter)
	logger.Out = os.Stdout

	//setup router
	router := mux.NewRouter()

	//setup the database
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, DBPORT, USER, PASSWORD, DBNAME)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return errors.Wrap(err, "Could not connect to the db")
	}
	defer db.Close()
	// This step is needed because db.Open() simply validates the arguments, it does not open an actual connection to the db.
	err = db.Ping()
	if err != nil {
		return err
	}
	err = setupDatabase(db)
	if err != nil {
		return errors.Wrap(err, "Could not setup the db")
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

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
