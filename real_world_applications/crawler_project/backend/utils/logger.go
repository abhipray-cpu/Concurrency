package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// CustomFormatter extends logrus to add your own prefix, for example
type CustomFormatter struct {
	logrus.Formatter
}

// Format an entry
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Message = "[MyApp] " + entry.Message
	return f.Formatter.Format(entry)
}

func NewLogger() *logrus.Logger {
	// Create a new instance of the logger
	logger := logrus.New()

	// Set the output (stdout, file, etc.)
	logger.Out = os.Stdout

	// Set the log level
	logger.SetLevel(logrus.InfoLevel)

	// Set formatter
	logger.SetFormatter(&CustomFormatter{
		&logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true, // Enable colors in the terminal
		},
	})

	// Example logging
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// If you want to log to a file instead of stdout
	file, err := os.OpenFile("../../log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	return logger
}
