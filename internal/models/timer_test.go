package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func getMockTimer() (*Timer, error) {
	id, err := uuid.Parse("09ef702a-3475-4cf4-8503-ad915f9abb5f")
	if err != nil {
		return nil, err
	}

	timer := Timer{
		ID:             id,
		Account:        "Hello",
		ExecutionCount: 10,
		taskConfig: &HTTP{
			URL:     "http://www.example.com/",
			Method:  GET,
			Body:    "hello go",
			Headers: []string{"set-cookie:yum"},
		},
		scheduleConfig: &Cron{
			Start:      time.Unix(0, 0).UTC(), // important, we're not encoding location so make sure it's set to nil (UTC)
			Cron:       "* * * * *",
			Executions: 10,
		},
		Meta: Meta{
			creationTime: time.Unix(0, 0).UTC(),
		},
	}

	return &timer, nil
}

func TestMarshalTimer(t *testing.T) {
	timer, err := getMockTimer()

	if err != nil {
		t.Errorf("Mock timer error: %v\n", err)
	}

	bytes, err := timer.MarshalBinary()
	if err != nil {
		t.Fail()
	}
	t.Logf("size of encoding: %v\n", len(bytes))

	var timer2 Timer
	err = timer2.UnmarshalBinary(bytes)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	if !reflect.DeepEqual(*timer, timer2) {
		t.Error("Unmarshaled timer unequal to original")
	}
}
