package repository

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) (*TaskRepository, error) {
	if db == nil {
		return &TaskRepository{}, errors.New("db must not be nil")
	}

	return &TaskRepository{db}, nil
}

func (r *TaskRepository) Add(ctx context.Context, task domain.Task) (domain.Task, error) {
	return domain.Task{}, nil
}

func (r *TaskRepository) List(ctx context.Context) ([]domain.Task, error) {
	return []domain.Task{}, nil
}

func (r *TaskRepository) ListByIDAndUserID(ctx context.Context, ID, userID int) (domain.Task, error) {
	return domain.Task{}, nil
}

func (r *TaskRepository) ListByUserID(ctx context.Context, userID int) ([]domain.Task, error) {
	return []domain.Task{}, nil
}
