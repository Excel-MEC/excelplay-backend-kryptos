package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
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
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stdout)

	//setup router
	router := mux.NewRouter()

	//setup the database
	db, err := setupDatabase()
	if err != nil {
		return errors.Wrap(err, "Could not setup the db")
	}

	server := &http.Server{
		Handler:      newServer(router, db),
		Addr:         PORT,
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	//start server
	logrus.Info("Server starting on port " + PORT)
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
