package main

import (
	"log"
	"log/syslog"
)

/*
	This program writes to syslog, not all platforms support this.
	You might need to modify the syslog configuration.
*/

var logger *log.Logger

func init() {
	var err error
	logger, err = syslog.NewLogger(syslog.LOG_USER|syslog.LOG_NOTICE, 0)
	if err != nil {
		log.Fatal("cannot write to syslog:", err)
	}
}

func main() {
	logger.Print("Hello syslog!!!!!!!!!")
}
