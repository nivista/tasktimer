package producers

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

type Timer struct {
}

func (*Timer) AddOrUpdate(t *models.Timer) error
func (*Timer) Delete(id uuid.UUID) error
