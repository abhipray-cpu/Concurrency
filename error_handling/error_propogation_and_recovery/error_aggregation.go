package main

import (
	"errors"
	"strings"
	"sync"
)

type ErrorAggregator struct {
	mu     sync.Mutex
	errors []error
}

func (e *ErrorAggregator) Add(err error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if err != nil {
		e.errors = append(e.errors, err)
	}
}

func (e *ErrorAggregator) Aggregate() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if len(e.errors) == 0 {
		return nil
	}
	var errMsgs []string
	for _, err := range e.errors {
		errMsgs = append(errMsgs, err.Error())
	}
	return errors.New(strings.Join(errMsgs, ", "))
}
