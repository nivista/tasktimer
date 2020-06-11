package models

import (
	"errors"
	"time"

	v1 "github.com/nivista/tasktimer/api/v1"

	"github.com/golang/protobuf/ptypes"
)

// Schedule is an object handled by a ScheduleVisitor
type Schedule interface {
	Visit(ScheduleVisitor)
}

// ScheduleVisitor should be implemented by anyone that wants to handle every kind of Schedule
type ScheduleVisitor interface {
	VisitCron(Cron)
	VisitInterval(Interval)
}

// Cron represents the configuration to a cron schedule
type Cron struct {
	Start      time.Time
	Cron       string
	Executions int32
}

// Visit calls VisitCron on the ScheduleVisitor
func (c Cron) Visit(v ScheduleVisitor) {
	v.VisitCron(c)
}

// Interval represents the configuration to an interval based scheduler
type Interval struct {
	Start      time.Time
	Interval   int32
	Executions int32
}

// Visit calls VisitInterval on the ScheduleVisitor
func (i Interval) Visit(v ScheduleVisitor) {
	v.VisitInterval(i)
}

type protoScheduleGenerator struct {
	*v1.Timer
	error
}

func (p protoScheduleGenerator) VisitCron(c Cron) {
	start, err := ptypes.TimestampProto(c.Start)
	p.error = err

	p.ScheduleConfig = &v1.Timer_CronConfig{CronConfig: &v1.CronConfig{
		StartTime:  start,
		Cron:       c.Cron,
		Executions: c.Executions,
	}}
}

func (p protoScheduleGenerator) VisitInterval(i Interval) {
	start, err := ptypes.TimestampProto(i.Start)
	p.error = err

	p.ScheduleConfig = &v1.Timer_IntervalConfig{IntervalConfig: &v1.IntervalConfig{
		StartTime:  start,
		Interval:   i.Interval,
		Executions: i.Executions,
	}}
}

func toScheduleConfig(p *v1.Timer) (Schedule, error) {
	switch config := p.ScheduleConfig.(type) {
	case *v1.Timer_CronConfig:
		pCronConfig := config.CronConfig

		start, err := ptypes.Timestamp(pCronConfig.StartTime)
		if err != nil {
			return nil, err
		}
		cronConfig := Cron{
			Start:      start,
			Cron:       pCronConfig.Cron,
			Executions: pCronConfig.Executions,
		}
		return &cronConfig, nil
	case *v1.Timer_IntervalConfig:
		pIntervalConfig := config.IntervalConfig

		start, err := ptypes.Timestamp(pIntervalConfig.StartTime)
		if err != nil {
			return nil, err
		}

		intervalConfig := Interval{
			Start:      start,
			Interval:   pIntervalConfig.Interval,
			Executions: pIntervalConfig.Executions,
		}
		return &intervalConfig, nil
	default:
		return nil, errors.New("Unable to parse")
	}
}
