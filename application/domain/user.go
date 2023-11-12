//go:generate mockgen -source=user.go -destination=user_mock.go -package=domain

package domain

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserNotAllowed  = errors.New("forbidden")
	ErrUserInvalidPass = errors.New("invalid password")
)

const (
	Manager    = "manager"
	Technician = "technician"
)

type User struct {
	ID       int64
	Email    string
	Password string
	RoleID   int64
}

type UserRetriever interface {
	ListByEmail(ctx context.Context, email string) (User, error)
}

func (u *User) GetRole() string {
	switch u.RoleID {
	case 1:
		return Manager
	case 2:
		return Technician
	}

	return ""
}
