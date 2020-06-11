package scheduler

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

// Manager manages recurring tasks
type Manager struct {
	timers map[uint32]struct {
		isExecuting bool
		timers      map[uuid.UUID]struct {
			stop  chan<- int
			timer *models.Timer
		}
	}

	db TimerEventListener
}

// TimerEventListener listens to changes in timer execution count.
type TimerEventListener interface {
	AddOrUpdate(*models.Timer) error
	Delete(*models.Timer) error
}

// NewManager creates a new manager.
func NewManager() (Manager, error)

// ExecutePartition starts the execution of all timers in memory for a partition.
func (*Manager) ExecutePartition(partition int) error

// DropPartition stops execution of, and removes from memory, all timers in a partition.
func (*Manager) DropPartition(partition int) error

// AddTimer adds a timer to a partition, or updates a timers execution count if it already exists.
func (*Manager) AddTimer(t *models.Timer, partition int) error

// RemoveTimer stops a timer and removes it.
func (*Manager) RemoveTimer(id uuid.UUID, partition int) error
