package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "excelplay-backend: ", log.LstdFlags)
}

// Println prints the given data on a new line
func Println(v interface{}) {
	logger.Println(v)
}

//Fatalln prints the given data and exits
func Fatalln(v interface{}) {
	logger.Fatalln(v)
}
