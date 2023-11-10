//go:generate mockgen -source=user.go -destination=user_mock.go -package=domain

package domain

import (
	"context"
)

const (
	Manager    = "manager"
	Technician = "technician"
)

type User struct {
	ID     int
	RoleID int
}

type UserRetriever interface {
	ListByUserID(ctx context.Context, userID int) (User, error)
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
