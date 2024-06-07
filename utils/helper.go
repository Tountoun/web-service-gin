package utils

import (
	"log"
	"os"
)


var Info *log.Logger

var Error *log.Logger


func init() {
	file, err := os.OpenFile("logging.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	Info = log.New(file, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
}