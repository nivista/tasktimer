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
		executable: &HTTPConfig{
			url:     "http://www.example.com/",
			method:  get,
			body:    "hello go",
			headers: []string{"set-cookie:yum"},
		},
		schedulable: &CronConfig{
			start:      time.Unix(0, 0).UTC(), // important, we're not encoding location so make sure it's set to nil (UTC)
			cron:       "* * * * *",
			executions: 10,
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

func TestCronSchedule(t *testing.T) {
	crons := []string{"0 0 12 1 * *"}
	cronConfigs := make([]CronConfig, len(crons))
	for idx, cron := range crons {
		cronConfigs[idx] = CronConfig{
			start: time.Unix(0, 0).UTC(),
			cron:  cron,
		}
	}

	now := time.Unix(0, 0).UTC().Add(time.Duration(1))

	expectedResults := []time.Time{time.Unix(60*60*12, 0).UTC()}
	for idx, config := range cronConfigs {
		timer, err := config.GetNextExec(&now)
		if err != nil {
			t.Errorf("failed test case %v with error: %v\n", idx, err)
		} else if timer != &expectedResults[idx] {
			t.Errorf("failed test case %v with unexpected result\n", idx)
		}
	}
}

func TestIntervalSchedule(t *testing.T) {
	intervals := []int32{1, 60, 1000}

	intervalConfigs := make([]IntervalConfig, len(intervals))

	for idx, interval := range intervals {
		intervalConfigs[idx] = IntervalConfig{
			start:    time.Unix(0, 0).UTC(),
			interval: interval,
		}
	}

	now := time.Unix(0, 0).UTC().Add(time.Duration(1))
	expectedResults := []time.Time{time.Unix(1, 0).UTC(),
		time.Unix(60, 0).UTC(),
		time.Unix(1000, 0).UTC(),
	}

	for idx, config := range intervalConfigs {
		timer, err := config.GetNextExec(&now)
		if err != nil {
			t.Errorf("failed test case %v with error: %v\n", idx, err)
		} else if timer != &expectedResults[idx] {
			t.Errorf("failed test case %v with unexpected result\n", idx)
		}
	}
}
