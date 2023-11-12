package domain

import "context"

type AuthUsecase interface {
	Authenticate(ctx context.Context, email, password string) (string, error)
}
