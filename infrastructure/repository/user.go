package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) (*UserRepository, error) {
	if db == nil {
		return &UserRepository{}, errors.New("db must not be nil")
	}

	return &UserRepository{db}, nil
}

func (r *UserRepository) ListByID(ctx context.Context, id int) (domain.User, error) {
	var result domain.User

	err := r.db.QueryRowContext(ctx, `SELECT id, role_id FROM users WHERE id=? AND deleted=FALSE`, id).Scan(&result.ID, &result.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, domain.ErrUserNotFound
		}

		return result, err
	}

	return result, nil
}

func (r *UserRepository) ListByEmail(ctx context.Context, email string) (domain.User, error) {
	var result domain.User

	err := r.db.QueryRowContext(ctx, `SELECT id, email, password, role_id FROM users WHERE email=? AND deleted=FALSE`, email).Scan(&result.ID, &result.Email, &result.Password, &result.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, domain.ErrUserNotFound
		}

		return result, err
	}

	return result, nil
}
