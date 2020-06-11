package consumers

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

func InitMemory(m manager) {

}

// Manager handles the management of tasks. Implementations should be safe for concurrent access.
type manager interface {
	ExecutePartition(partition int) error
	DropPartition(partition int) error
	AddTimer(t *models.Timer, partition int) error
	RemoveTimer(id uuid.UUID, partition int) error
}
