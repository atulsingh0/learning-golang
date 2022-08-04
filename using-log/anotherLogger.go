package main

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
)

func main() {

	infoLogger.Println("this is my first log msg.")
	warnLogger.Println("this is warning log msg.")
	errLogger.Println("and last, error log msg.")

}

func init() {

	//  init func will auto execute

	logFile := "anotherLogger.log"
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0700)

	if err != nil {
		log.Fatal("Unable to create log file.", err)
	}
	log.SetOutput(file)

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	fatalLogger = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}
