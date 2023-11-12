//go:generate mockgen -source=user.go -destination=user_mock.go -package=domain

package domain

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

const (
	Manager    = "manager"
	Technician = "technician"
)

type User struct {
	ID       int
	Email    string
	Password string
	RoleID   int
}

type UserRetriever interface {
	ListByID(ctx context.Context, userID int) (User, error)
	ListByEmail(ctx context.Context, email string) (User, error)
}

type UserAuthenticator interface {
	GenerateAccessToken(User) (string, error)
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
