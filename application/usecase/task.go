package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"log"
	"time"
)

type taskUseCase struct {
	creator   domain.TaskCreator
	retriever domain.TaskRetriever
	updater   domain.TaskUpdater
	remover   domain.TaskRemover
	encryptor domain.SummaryEncryptor
}

func NewTask(
	creator domain.TaskCreator,
	retriever domain.TaskRetriever,
	updater domain.TaskUpdater,
	remover domain.TaskRemover,
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

	if encryptor == nil {
		return &taskUseCase{}, errors.New("encryptor must not be nil")
	}

	return &taskUseCase{
		creator,
		retriever,
		updater,
		remover,
		encryptor,
	}, nil
}

func (u *taskUseCase) ListByUser(ctx context.Context, user domain.User) ([]domain.Task, error) {
	if user.ID == 0 {
		return []domain.Task{}, errors.New("user ID must not be empty")
	}

	if user.RoleID == 0 {
		return []domain.Task{}, errors.New("user RoleID must not be empty")
	}

	var err error
	var tasks []domain.Task

	if user.GetRole() == domain.Manager {
		tasks, err = u.retriever.List(ctx)
	}

	if user.GetRole() == domain.Technician {
		tasks, err = u.retriever.ListByUserID(ctx, user.ID)
	}

	if err != nil {
		return []domain.Task{}, err
	}

	var decryptedTasks []domain.Task

	for _, t := range tasks {
		summaryDecrypt, err := u.encryptor.Decrypt(t.Summary)
		if err != nil {
			log.Printf("Error decrypting task summary: %v", err) // Later: send to metrics/observability
			continue
		}

		t.Summary = summaryDecrypt
		decryptedTasks = append(decryptedTasks, t)
	}

	return decryptedTasks, nil
}

func (u *taskUseCase) Add(ctx context.Context, task domain.Task, user domain.User) (domain.Task, error) {
	if task.UserID == 0 {
		return domain.Task{}, errors.New("UserID must not be 0")
	}

	if user.ID == 0 {
		return domain.Task{}, errors.New("user ID must not be empty")
	}

	if user.RoleID == 0 {
		return domain.Task{}, errors.New("user RoleID must not be empty")
	}

	if user.ID != task.UserID {
		return domain.Task{}, errors.New("forbidden")
	}

	summaryEncrypted, err := u.encryptor.Encrypt(task.Summary)
	if err != nil {
		return domain.Task{}, err
	}
	task.Summary = summaryEncrypted

	id, err := u.creator.Add(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	task.ID = id

	summaryDecrypt, err := u.encryptor.Decrypt(task.Summary)
	if err != nil {
		return domain.Task{}, err
	}

	task.Summary = summaryDecrypt

	return task, nil
}

func (u *taskUseCase) Update(ctx context.Context, task domain.Task, user domain.User) (domain.Task, error) {
	if task.ID == 0 {
		return domain.Task{}, errors.New("ID must not be 0")
	}

	if task.UserID == 0 {
		return domain.Task{}, errors.New("UserID must not be 0")
	}

	if user.ID == 0 {
		return domain.Task{}, errors.New("user ID must not be empty")
	}

	if user.RoleID == 0 {
		return domain.Task{}, errors.New("user RoleID must not be empty")
	}

	if user.ID != task.UserID {
		return domain.Task{}, errors.New("forbidden")
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

	err = u.updater.Update(ctx, tsk)
	if err != nil {
		return domain.Task{}, err
	}

	summaryDecrypt, err := u.encryptor.Decrypt(tsk.Summary)
	if err != nil {
		return domain.Task{}, err
	}

	tsk.Summary = summaryDecrypt

	return tsk, nil
}

func (u *taskUseCase) Remove(ctx context.Context, id int, user domain.User) error {
	if id == 0 {
		return errors.New("ID must not be 0")
	}

	if user.ID == 0 {
		return errors.New("user ID must not be empty")
	}

	if user.RoleID == 0 {
		return errors.New("user RoleID must not be empty")
	}

	if user.GetRole() == domain.Technician {
		return errors.New("forbidden")
	}

	return u.remover.Remove(ctx, id)
}
