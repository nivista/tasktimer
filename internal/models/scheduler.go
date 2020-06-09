package models

import (
	"errors"
	"time"

	v1 "github.com/nivista/tasktimer/api/v1"

	"github.com/golang/protobuf/ptypes"
)

type schedulable interface {
	assignToProto(*v1.Timer) error
	GetNextExec(*time.Time) (*time.Time, error)
	IsValid() (bool, string)
}

// CronConfig represents a scheduling scheme based on a 6 digit crontab.
// "* * * * * *", like regular cron but seconds is the first field.
type CronConfig struct {
	start      time.Time
	cron       string
	executions int32
}

// IntervalConfig represents a scheduling scheme in which the timer fires every
// "interval" seconds
type IntervalConfig struct {
	start      time.Time
	interval   int32
	executions int32
}

func (c *CronConfig) assignToProto(p *v1.Timer) error {
	start, err := ptypes.TimestampProto(c.start)
	if err != nil {
		return err
	}
	p.SchedulerConfig = &v1.Timer_CronConfig{CronConfig: &v1.CronConfig{
		StartTime:  start,
		Cron:       c.cron,
		Executions: c.executions,
	}}
	return nil
}

// GetNextExec returns a time.Timer that will fire next time an execution should happen
func (c *CronConfig) GetNextExec(*time.Time) (*time.Time, error) {
	return nil, nil
}

// IsValid returns true, "" if the configuration is valid. Otherwise it returns false, and a reason.
func (c *CronConfig) IsValid() (bool, string) {
	return true, ""
}

func (c *IntervalConfig) assignToProto(p *v1.Timer) error {
	start, err := ptypes.TimestampProto(c.start)
	if err != nil {
		return err
	}
	p.SchedulerConfig = &v1.Timer_IntervalConfig{IntervalConfig: &v1.IntervalConfig{
		StartTime:  start,
		Interval:   c.interval,
		Executions: c.executions,
	}}
	return nil
}

// GetNextExec returns a time.Timer that will fire next time an execution should happen
func (c *IntervalConfig) GetNextExec(*time.Time) (*time.Time, error) {
	return nil, nil
}

// IsValid returns true, "" if the configuration is valid. Otherwise it returns false, and a reason.
func (c *IntervalConfig) IsValid() (bool, string) {
	return true, ""
}

func toSchedulable(p *v1.Timer) (schedulable, error) {
	switch config := p.SchedulerConfig.(type) {
	case *v1.Timer_CronConfig:
		pCronConfig := config.CronConfig

		start, err := ptypes.Timestamp(pCronConfig.StartTime)
		if err != nil {
			return nil, err
		}
		cronConfig := CronConfig{
			start:      start,
			cron:       pCronConfig.Cron,
			executions: pCronConfig.Executions,
		}
		return &cronConfig, nil
	case *v1.Timer_IntervalConfig:
		pIntervalConfig := config.IntervalConfig

		start, err := ptypes.Timestamp(pIntervalConfig.StartTime)
		if err != nil {
			return nil, err
		}

		intervalConfig := IntervalConfig{
			start:      start,
			interval:   pIntervalConfig.Interval,
			executions: pIntervalConfig.Executions,
		}
		return &intervalConfig, nil
	default:
		return nil, errors.New("Unable to parse")
	}
}
