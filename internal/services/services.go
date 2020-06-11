package services

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

type dataAccess interface {
	GetTimer(account string, id uuid.UUID) (models.Timer, error)
	GetTimers(account string) ([]*models.Timer, error)
}
