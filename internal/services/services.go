package services

import (
	"time"

	"github.com/google/uuid"
	v1 "github.com/nivista/tasktimer/api/v1"
	"github.com/nivista/tasktimer/internal/models"
)

type dataAccess interface {
	GetTimer(account string, id uuid.UUID) (models.Timer, error)
	GetTimers(account string) ([]uuid.UUID, error)
}

type timerProd interface {
	AddOrUpdate(t *models.Timer) error
	Delete(id uuid.UUID) error
}

// Service handles external interaction
type Service struct {
	db   dataAccess
	prod timerProd
}

// InitServices returns an object that handles the external interaction
func InitServices(db dataAccess, prod timerProd) *Service {
	return &Service{
		db:   db,
		prod: prod,
	}
}

// AddTimer adds a timer.
func (s *Service) AddTimer(c *v1.CreateTimer) error {
	task, err := models.ToTask(c.TaskConfig)
	if err != nil {
		return err
	}
	schedule, err := models.ToSchedule(c.ScheduleConfig)
	if err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	timer := models.Timer{
		ID:      id,
		Account: c.Account,
		Meta: models.Meta{
			CreationTime: time.Now(),
		},
		Task:     task,
		Schedule: schedule,
	}

	// TODO: authentication and validation

	s.prod.AddOrUpdate(&timer)

	return nil
}

// RemoveTimer removes a timer
func (s *Service) RemoveTimer(account string, id uuid.UUID) error {
	return s.prod.Delete(id)
}

// GetTimer gets a timer
func (s *Service) GetTimer(account string, id uuid.UUID) (models.Timer, error) {
	return s.db.GetTimer(account, id)
}

// GetTimers gets ids of all timers associated with an account
func (s *Service) GetTimers(account string) ([]uuid.UUID, error) {
	return s.db.GetTimers(account)
}
