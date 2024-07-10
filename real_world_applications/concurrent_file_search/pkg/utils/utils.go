package utils

import (
	"fmt"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

// LogMessage logs a message with a given severity level.
func LogMessage(level int, message string) {
	prefix := ""
	switch level {
	case DEBUG:
		prefix = "DEBUG"
	case INFO:
		prefix = "INFO"
	case WARN:
		prefix = "WARN"
	case ERROR:
		prefix = "ERROR"
	}

	// Format the log message with a timestamp
	logMessage := fmt.Sprintf("%s [%s] %s", time.Now().Format(time.RFC3339), prefix, message)
	fmt.Println(logMessage)
}
