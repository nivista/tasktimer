package models

import (
	"time"

	v1 "github.com/nivista/tasktimer/api/v1"

	"github.com/golang/protobuf/ptypes"
)

// Meta represents metadata about a timer
type Meta struct {
	CreationTime time.Time
}

func (m *Meta) assignToProto(p *v1.Timer) error {
	pTime, err := ptypes.TimestampProto(m.CreationTime)
	if err != nil {
		return err
	}
	p.Meta = &v1.Meta{
		CreateTime: pTime,
	}
	return nil
}

func getMeta(p *v1.Timer) (*Meta, error) {
	creationTime, err := ptypes.Timestamp(p.Meta.CreateTime)
	if err != nil {
		return nil, err
	}
	return &Meta{
		CreationTime: creationTime,
	}, nil
}
