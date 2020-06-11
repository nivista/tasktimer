package access

import (
	"github.com/google/uuid"
	"github.com/nivista/tasktimer/internal/models"
)

// DB represents a database.
type DB struct {
}

// InsertTimer inserts a timer into DB.
func (d *DB) InsertTimer(t *models.Timer) error

// DeleteTimer deletes a timer from the DB.
func (d *DB) DeleteTimer(account string, id uuid.UUID) error

// SetExecCount sets the execution count of a timer in the DB.
func (d *DB) SetExecCount(account string, id uuid.UUID, count int) error

// GetTimer gets a timer from the DB.
func (d *DB) GetTimer(account string, id uuid.UUID) (models.Timer, error)

// GetTimers gets all timers associated with an account from the DB.
func (d *DB) GetTimers(account string) ([]uuid.UUID, error)
