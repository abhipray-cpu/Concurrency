// logger.go
package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type Logger struct {
	logChan    chan string
	entryIndex int64
}

func NewLogger(bufferSize int) *Logger {
	return &Logger{
		logChan:    make(chan string, bufferSize),
		entryIndex: 0,
	}
}

func (l *Logger) LogMessage(message string) {
	index := atomic.AddInt64(&l.entryIndex, 1)
	l.logChan <- fmt.Sprintf("%d: %s", index, message)
}

func (l *Logger) StartLogWriter(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for msg := range l.logChan {
		_, err := file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " - " + msg + "\n")
		if err != nil {
			panic(err)
		}
	}
}
