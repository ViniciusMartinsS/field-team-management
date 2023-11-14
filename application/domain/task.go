//go:generate mockgen -source=task.go -destination=task_mock.go -package=domain

package domain

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrInvalidTask   = errors.New("task fields are invalid")
	ErrTasksNotFound = errors.New("tasks not found")
	TaskDateLayout   = "01/02/2006 15:04"
)

type TaskUsecase interface {
	Add(ctx context.Context, task Task, user User) (Task, error)
	ListByUser(ctx context.Context, user User) ([]Task, error)
	Update(ctx context.Context, task Task, user User) (Task, error)
	Remove(ctx context.Context, id int64, user User) error
}

type TaskCreator interface {
	Add(ctx context.Context, task Task) (int64, error)
}

type TaskRetriever interface {
	List(ctx context.Context) ([]Task, error)
	ListByIDAndUserID(ctx context.Context, id, userID int64) (Task, error)
	ListByUserID(ctx context.Context, userID int64) ([]Task, error)
}

type TaskUpdater interface {
	Update(ctx context.Context, task Task) error
}

type TaskRemover interface {
	Remove(ctx context.Context, id int64) error
}

type TaskNotifier interface {
	SendNotification(ctx context.Context, body string) error
}

type SummaryEncryptor interface {
	Encrypt(value string) (string, error)
	Decrypt(value string) (string, error)
}

type Task struct {
	ID      int64
	Summary string
	Date    *time.Time
	UserID  int64
}

func NewTask(summary string, date *time.Time, userID int64) (Task, error) {
	var err []string

	if summary == "" {
		err = append(err, "summary must not be empty")
	}

	if len(summary) > 2500 {
		err = append(err, "summary size is too big")
	}

	if userID == 0 {
		err = append(err, "user ID must not be 0")
	}

	if len(err) > 0 {
		return Task{}, fmt.Errorf("%v: %s", ErrInvalidTask, strings.Join(err, "; "))
	}

	return Task{Summary: summary, Date: date, UserID: userID}, nil
}
