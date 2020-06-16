package timer

import (
	"encoding"
	"encoding/json"
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
		encoding.BinaryMarshaler
		encoding.BinaryUnmarshaler
	}

	// Schedule represents the "when" configuration of a timer.
	Schedule interface {
		isSchedule()
		encoding.BinaryMarshaler
		encoding.BinaryUnmarshaler
	}

	Meta struct {
		CreationTime time.Time
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

func (t *HTTP) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *HTTP) UnmarshalBinary(bytes []byte) error {
	return json.Unmarshal(bytes, t)
}

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

func (s *Cron) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Cron) UnmarshalBinary(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}

func (Interval) isSchedule() {}

func (s *Interval) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Interval) UnmarshalBinary(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
