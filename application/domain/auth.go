//go:generate mockgen -source=auth.go -destination=auth_mock.go -package=domain

package domain

import "context"

type AuthUsecase interface {
	Authenticate(ctx context.Context, email, password string) (string, error)
}

type Authenticator interface {
	GenerateAccessToken(User) (string, error)
	IsAccessTokenValid(string) (bool, map[string]any, error)
}
