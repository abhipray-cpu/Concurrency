package main

// Job represents a task that can be executed.
type Job struct {
	ID       string
	Function func() error
}

// NewJob creates a new Job instance.
func NewJob(id string, function func() error) *Job {
	return &Job{
		ID:       id,
		Function: function,
	}
}
