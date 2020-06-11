package services

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

type dataAccess interface {
	GetTimer(account string, id uuid.UUID) (models.Timer, error)
	GetTimers(account string) ([]*models.Timer, error)
}

type timerProd interface {
	Add()
	Remove()
}

// Service handles external interaction
type Service struct {
	dataAccess
	timerProd
}

// InitServices returns an object that handles the external interaction
func InitServices(db dataAccess, prod timerProd) *Service {
	return &Service{
		dataAccess: db,
		timerProd:  prod,
	}
}
