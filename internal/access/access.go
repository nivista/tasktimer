package access

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

// DB is the interface to the database
type DB interface {
	InsertTimer(t *models.Timer) error
	DeleteTimer(account string, id uuid.UUID) error
	GetTimer(account string, id uuid.UUID) (models.Timer, error)
	GetTimers(account string) ([]*models.Timer, error)
	SetExecCount(account string, id uuid.UUID, count int) error
}
