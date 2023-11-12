package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	authenticator domain.Authenticator
	retriever     domain.UserRetriever
}

func NewAuth(authenticator domain.Authenticator, retriever domain.UserRetriever) (domain.AuthUsecase, error) {
	if authenticator == nil {
		return &authUseCase{}, errors.New("authenticator must not be nil")
	}

	if retriever == nil {
		return &authUseCase{}, errors.New("retriever must not be nil")
	}

	return &authUseCase{
		authenticator: authenticator,
		retriever:     retriever,
	}, nil
}

func (u *authUseCase) Authenticate(ctx context.Context, email, password string) (string, error) {
	user, err := u.retriever.ListByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", domain.ErrUserInvalidPass
	}

	token, err := u.authenticator.GenerateAccessToken(user)
	if err != nil {
		return "", errors.New("error while generating token")
	}

	return token, nil
}
