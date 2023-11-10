package domain

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrInvalidTask = errors.New("task fields are invalid")

type TaskCreator interface {
	Add(ctx context.Context, task Task) (Task, error)
}

type Task struct {
	ID      int
	Summary string
	Date    *time.Time
	UserID  int
}

func NewTask(summary string, date *time.Time, userID int) (Task, error) {
	var err []string

	if summary == "" {
		err = append(err, "summary must not be empty")
	}

	if userID == 0 {
		err = append(err, "user ID must not be 0")
	}

	if len(err) > 0 {
		return Task{}, fmt.Errorf("%v: %s", ErrInvalidTask, strings.Join(err, "; "))
	}

	return Task{Summary: summary, Date: date, UserID: userID}, nil
}
