package timer

import (
	"time"

	"github.com/google/uuid"
)

// Timer definition
type (
	Timer struct {
		ID             uuid.UUID
		Account        string
		ExecutionCount uint32
		Task
		Schedule
		Meta
	}

	// Task represents a timer task.
	Task interface {
		isTask()
	}

	// Schedule represents the "when" configuration of a timer.
	Schedule interface {
		isSchedule()
	}
)

// HTTP Task represents an HTTP request.
type (
	HTTP struct {
		URL string
		Method
		Body    string
		Headers []string
	}

	// Method is the method of the HTTP request.
	Method int
)

func (HTTP) isTask() {}

// Method definitions
const (
	GET Method = iota
	POST
)

// Schedule definitions
type (
	// Cron is a schedule configured by a cron string.
	Cron struct {
		Start      time.Time
		Cron       string
		Executions int32
	}

	// Interval is a schedule configured by a fixed interval.
	Interval struct {
		Start      time.Time
		Interval   int32
		Executions int32
	}
)

func (Cron) isSchedule() {}

func (Interval) isSchedule() {}
