package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"time"
)

type taskUseCase struct {
	creator       domain.TaskCreator
	retriever     domain.TaskRetriever
	updater       domain.TaskUpdater
	remover       domain.TaskRemover
	userRetriever domain.UserRetriever
	encryptor     domain.SummaryEncryptor
}

func NewTask(
	creator domain.TaskCreator,
	retriever domain.TaskRetriever,
	updater domain.TaskUpdater,
	remover domain.TaskRemover,
	userRetriever domain.UserRetriever,
	encryptor domain.SummaryEncryptor,
) (domain.TaskUsecase, error) {
	if creator == nil {
		return &taskUseCase{}, errors.New("task creator must not be nil")
	}

	if retriever == nil {
		return &taskUseCase{}, errors.New("task retriever must not be nil")
	}

	if updater == nil {
		return &taskUseCase{}, errors.New("task updater must not be nil")
	}

	if remover == nil {
		return &taskUseCase{}, errors.New("task remover must not be nil")
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
		updater,
		remover,
		userRetriever,
		encryptor,
	}, nil
}

func (u *taskUseCase) ListByUserID(ctx context.Context, userID int) ([]domain.Task, error) {
	user, err := u.userRetriever.ListByID(ctx, userID)
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

func (u *taskUseCase) Update(ctx context.Context, task domain.Task) (domain.Task, error) {
	if task.ID == 0 {
		return domain.Task{}, errors.New("ID must not be 0")
	}

	if task.UserID == 0 {
		return domain.Task{}, errors.New("UserID must not be 0")
	}

	tsk, err := u.retriever.ListByIDAndUserID(ctx, task.ID, task.UserID)
	if err != nil {
		return domain.Task{}, err
	}

	var cp *time.Time
	if task.Date != cp {
		tsk.Date = task.Date
	}

	if task.Summary != "" {
		summaryEncrypted, err := u.encryptor.Encrypt(task.Summary)
		if err != nil {
			return domain.Task{}, err
		}

		tsk.Summary = summaryEncrypted
	}

	t, err := u.updater.Update(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return t, nil
}

func (u *taskUseCase) Remove(ctx context.Context, id, userID int) error {
	if id == 0 {
		return errors.New("ID must not be 0")
	}

	if userID == 0 {
		return errors.New("UserID must not be 0")
	}

	user, err := u.userRetriever.ListByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.GetRole() == domain.Technician {
		return errors.New("forbidden")
	}

	return u.remover.Remove(ctx, id, userID)
}
