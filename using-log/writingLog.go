package main

import (
	"log"
	"os"
)

type msgType int

const (
	INFO msgType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func main() {

	logFile := "writingLog.log"

	writeLog(logFile, INFO, "this is my first log msg.")
	writeLog(logFile, WARNING, "this is warning log msg.")
	writeLog(logFile, ERROR, "and last, error log msg.")

}

func writeLog(logFile string, msgType msgType, msg string) {

	// Opening Error file
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0700)

	if err != nil {
		log.Fatal("Unable to create log file.", err)
	}
	log.SetOutput(file)
	defer file.Close()

	switch msgType {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(msg)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(msg)
	}
}
