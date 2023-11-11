package repository

import (
	"context"
	"database/sql"
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
	rows, err := r.db.QueryContext(ctx, `SELECT id, summary, date, user_id FROM tasks WHERE deleted=FALSE`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Task{}, domain.ErrTasksNotFound
		}

		return []domain.Task{}, err
	}

	result, err := rowsToTask(rows)
	if err != nil {
		return []domain.Task{}, err
	}

	return result, nil
}

func (r *TaskRepository) ListByIDAndUserID(ctx context.Context, ID, userID int) (domain.Task, error) {
	return domain.Task{}, nil
}

func (r *TaskRepository) ListByUserID(ctx context.Context, userID int) ([]domain.Task, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, summary, date, user_id FROM tasks WHERE user_id=? AND deleted=FALSE`, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Task{}, domain.ErrTasksNotFound
		}

		return []domain.Task{}, err
	}

	result, err := rowsToTask(rows)
	if err != nil {
		return []domain.Task{}, err
	}

	return result, nil
}

func (r *TaskRepository) Update(ctx context.Context, task domain.Task) (domain.Task, error) {
	return domain.Task{}, nil
}

func (r *TaskRepository) Remove(ctx context.Context, id, userID int) error {
	return nil
}

func rowsToTask(rows *sql.Rows) ([]domain.Task, error) {
	var result []domain.Task

	for rows.Next() {
		var task domain.Task

		if err := rows.Scan(&task.ID, &task.Summary, &task.Date, &task.UserID); err != nil {
			return []domain.Task{}, errors.New("internal server error")
		}

		result = append(result, task)
	}

	return result, nil
}
