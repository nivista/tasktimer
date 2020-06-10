package scheduler

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

// Represents task scheduler. Implementations should be safe for concurrent access.
type scheduler interface {
	ExecutePartition(partition int) error
	DropPartition(partition int) error
	AddTimer(t *models.Timer, partition int) error
	RemoveTimer(id uuid.UUID, partition int) error
}
