package consumer

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

func InitAccess(db dataAccess) {

}

// DB is the interface to the database
type dataAccess interface {
	InsertTimer(t *models.Timer) error
	DeleteTimer(account string, id uuid.UUID) error
	SetExecCount(account string, id uuid.UUID, count int) error
}
