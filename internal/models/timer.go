package models

import (
	"github.com/google/uuid"
	v1 "github.com/nivista/tasktimer/api/v1"
	"google.golang.org/protobuf/proto"
)

// Timer represents the instructions and metadata for a recurring task
type Timer struct {
	ID             uuid.UUID
	Account        string
	ExecutionCount uint32
	Task
	Schedule
	Meta Meta // Not sure yet if this should be exported
}

// MarshalBinary returns a binary representation of a timer
func (t *Timer) MarshalBinary() ([]byte, error) {
	p, err := t.toProto()

	if err != nil {
		return nil, err
	}
	return proto.Marshal(p)
}

// UnmarshalBinary initializes a timer from a []byte created with MarshalBinary
func (t *Timer) UnmarshalBinary(blob []byte) error {
	var p v1.Timer
	if err := proto.Unmarshal(blob, &p); err != nil {
		return err
	}
	timer, err := fromProto(&p)
	if err != nil {
		return err
	}
	*t = *timer
	return nil
}

func fromProto(p *v1.Timer) (*Timer, error) {
	id, err := uuid.FromBytes(p.Id)
	if err != nil {
		return nil, err
	}

	taskConfig, err := toTaskConfig(p)
	if err != nil {
		return nil, err
	}

	scheduleConfig, err := toScheduleConfig(p)
	if err != nil {
		return nil, err
	}

	meta, err := getMeta(p)
	if err != nil {
		return nil, err
	}

	timer := Timer{
		ID:             id,
		Account:        p.Account,
		ExecutionCount: p.ExecutionCount,
		Task:           taskConfig,
		Schedule:       scheduleConfig,
		Meta:           *meta,
	}
	return &timer, nil
}

func (t *Timer) toProto() (*v1.Timer, error) {
	p := v1.Timer{}
	id, err := t.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	p.Id = id
	p.Account = t.Account
	p.ExecutionCount = t.ExecutionCount
	t.Task.Visit(protoTaskGenerator{&p})
	t.Schedule.Visit(protoScheduleGenerator{&p})
	t.Meta.assignToProto(&p)
	return &p, nil
}

/* const numPartitions = 1000 // TODO : env
func bytesToPartition(id []byte) int {
	partition := 0

	for b := range id {
		partition = partition << 8
		partition += b
		partition = partition % numPartitions
	}

	return partition
}
*/
