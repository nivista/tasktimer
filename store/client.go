package store

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nivista/tasktimer/timer"
)

type Client struct {
	*pgxpool.Pool
}

func NewClient() (*Client, error) {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &Client{Pool: dbpool}, nil
}

func (c *Client) CreateTimer(ctx context.Context, t timer.Timer) error {
	var taskType string
	switch task := t.Task.(type) {
	case timer.HTTP:
		taskType = "http"
	default:
		panic(fmt.Sprintf("CreateTimer doesn't know about %T", task))
	}

	var scheduleType string
	switch schedule := t.Schedule.(type) {
	case timer.Cron:
		scheduleType = "cron"
	case timer.Interval:
		scheduleType = "interval"
	default:
		panic(fmt.Sprintf("CreateTimer doesn't know about %T", schedule))
	}

	_, err := c.Query(ctx,
		`insert into timers (id, account, executionCount, taskType, task, scheduleType, schedule, meta) 
			values ($1, $2, $3, $4, $5, $6)
			where exists select id from timers where account=$2 and id=$1`,
		t.ID, t.Account, t.ExecutionCount, taskType, t.Task, scheduleType, t.Schedule, t.Meta)

	return err
}

func (c *Client) DeleteTimer(ctx context.Context, account string, id uuid.UUID) error {
	_, err := c.Query(ctx, //TODO check if a timer was deleted
		`delete from timers where account=$1 and id=$2`,
		account, id)
	return err
}

func (c *Client) GetTimer(ctx context.Context, account string, id uuid.UUID) (*timer.Timer, error) {
	t := timer.Timer{}

	var taskType, scheduleType string
	var task, schedule interface{}

	err := c.QueryRow(ctx,
		`select id, account, executionCount, taskType, task, scheduleType, schedule, meta where account=$1 and id=$2`,
		account, id).Scan(&t.ID, &t.Account, &t.ExecutionCount, &taskType, &task, &scheduleType, &schedule)

	if err != nil {
		return nil, err
	}

	switch taskType {
	case "http":

		http, ok := task.(timer.HTTP)
		if !ok {
			return nil, errors.New("Failed type declaration") // TODO, make custom error types
		}
		t.Task = http
	default:
		return nil, errors.New("Unknown type")
	}

	switch scheduleType {
	case "cron":
		cron, ok := task.(timer.Cron)
		if !ok {
			return nil, errors.New("Failed type declaration")
		}
		t.Schedule = cron
	case "interval":
		interval, ok := task.(timer.Interval)
		if !ok {
			return nil, errors.New("Failed type declaration")
		}
		t.Schedule = interval
	default:
		return nil, errors.New("Unknown type")
	}
	return &t, nil
}

func (c *Client) SetExecCount(ctx context.Context, account string, id uuid.UUID, count int) error {
	_, err := c.Query(ctx,
		`update timers
			set executionCount=$1
			where account=$2 and id=$3`,
		count, account, id)
	return err
}
