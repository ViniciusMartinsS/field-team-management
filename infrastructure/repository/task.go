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

func (r *TaskRepository) Add(ctx context.Context, task domain.Task) (int64, error) {
	var id int64

	record, err := r.db.ExecContext(ctx, `INSERT INTO tasks (summary, date, user_id) VALUES (?, ?, ?)`, task.Summary, task.Date, task.UserID)
	if err != nil {
		return id, err
	}

	if id, err = record.LastInsertId(); err != nil {
		return id, err
	}

	return id, nil
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

func (r *TaskRepository) ListByIDAndUserID(ctx context.Context, id, userID int64) (domain.Task, error) {
	var result domain.Task

	err := r.db.QueryRowContext(ctx, `SELECT id, summary, date, user_id FROM tasks WHERE id=? AND user_id=? AND deleted=FALSE`, id, userID).
		Scan(&result.ID, &result.Summary, &result.Date, &result.UserID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, domain.ErrTasksNotFound
		}

		return result, err
	}

	return result, nil
}

func (r *TaskRepository) ListByUserID(ctx context.Context, userID int64) ([]domain.Task, error) {
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

func (r *TaskRepository) Update(ctx context.Context, task domain.Task) error {
	_, err := r.db.ExecContext(ctx, `UPDATE tasks SET summary=?, date=? WHERE id=? AND user_id=? AND deleted=FALSE`, task.Summary, task.Date, task.ID, task.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrTasksNotFound
		}

		return err
	}

	return nil
}

func (r *TaskRepository) Remove(ctx context.Context, id int64) error {
	if _, err := r.db.ExecContext(ctx, `UPDATE tasks SET deleted=TRUE WHERE id=?`, id); err != nil {
		return err
	}

	return nil
}

func rowsToTask(rows *sql.Rows) ([]domain.Task, error) {
	var result []domain.Task

	for rows.Next() {
		var task domain.Task

		if err := rows.Scan(&task.ID, &task.Summary, &task.Date, &task.UserID); err != nil {
			return []domain.Task{}, err
		}

		result = append(result, task)
	}

	return result, nil
}
