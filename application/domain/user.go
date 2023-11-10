//go:generate mockgen -source=user.go -destination=user_mock.go -package=domain

package domain

import (
	"context"
)

type User struct {
	ID     int
	RoleID int
}

type UserRetriever interface {
	ListByUserID(ctx context.Context, userID int) (User, error)
}
