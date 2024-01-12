package config

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFileName string) {
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Faild to open log file: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.SetOutput(multiWriter)
}
