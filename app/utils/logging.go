package utils

import (
	"io"
	"log"
	"os"
)

func LoggingConfig(logPath string) {
	logFile, err := os.OpenFile(
		logPath,
		// readwrite, create if not exists, append if exists
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		// permission setting
		0666,
	)
	if err != nil {
		log.Fatalln(err)
	}

	// set log output to both stdout and logfile
	multiLogFile := io.MultiWriter(os.Stdout, logFile)
	// log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
