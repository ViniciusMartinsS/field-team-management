package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
)

type taskUseCase struct {
	creator       domain.TaskCreator
	retriever     domain.TaskRetriever
	userRetriever domain.UserRetriever
	encryptor     domain.SummaryEncryptor
}

func NewTask(creator domain.TaskCreator, retriever domain.TaskRetriever, userRetriever domain.UserRetriever, encryptor domain.SummaryEncryptor) (domain.TaskUsecase, error) {
	if creator == nil {
		return &taskUseCase{}, errors.New("task creator must not be nil")
	}

	if retriever == nil {
		return &taskUseCase{}, errors.New("task retriever must not be nil")
	}

	if userRetriever == nil {
		return &taskUseCase{}, errors.New("user retriever must not be nil")
	}

	if encryptor == nil {
		return &taskUseCase{}, errors.New("encryptor must not be nil")
	}

	return &taskUseCase{
		creator,
		retriever,
		userRetriever,
		encryptor,
	}, nil
}

func (u *taskUseCase) ListByUserID(ctx context.Context, userID int) ([]domain.Task, error) {
	user, err := u.userRetriever.ListByUserID(ctx, userID)
	if err != nil {
		return []domain.Task{}, err
	}

	var tasks []domain.Task

	if user.GetRole() == domain.Manager {
		tasks, err = u.retriever.List(ctx)
	}

	if user.GetRole() == domain.Technician {
		tasks, err = u.retriever.ListByUserID(ctx, userID)
	}

	if err != nil {
		return []domain.Task{}, err
	}

	return tasks, nil
}

func (u *taskUseCase) Add(ctx context.Context, task domain.Task) (domain.Task, error) {
	summaryEncrypted, err := u.encryptor.Encrypt(task.Summary)
	if err != nil {
		// Handle error to be more friendly
		return domain.Task{}, err
	}
	task.Summary = summaryEncrypted

	t, err := u.creator.Add(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return t, nil
}
