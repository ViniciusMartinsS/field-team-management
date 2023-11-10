package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"time"
)

type taskUseCase struct {
	creator   domain.TaskCreator
	retriever domain.TaskRetriever
}

func NewTask(creator domain.TaskCreator, retriever domain.TaskRetriever) (domain.TaskUsecase, error) {
	if creator == nil {
		return &taskUseCase{}, errors.New("task creator must not be nil")
	}

	if retriever == nil {
		return &taskUseCase{}, errors.New("task retriever must not be nil")
	}

	return &taskUseCase{
		creator,
		retriever,
	}, nil
}

func (u *taskUseCase) ListByUserID(ctx context.Context, userID int) ([]domain.Task, error) {
	currentTime := time.Now()
	return []domain.Task{
		{
			ID:      1,
			Summary: "Dummy Summary",
			Date:    &currentTime,
			UserID:  userID,
		},
	}, nil
}

func (u *taskUseCase) Add(ctx context.Context, task domain.Task) (domain.Task, error) {
	// I'll add logic to create the customer
	return task, nil
}
