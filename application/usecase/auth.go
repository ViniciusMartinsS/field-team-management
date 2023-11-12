package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	authenticator domain.UserAuthenticator
	retriever     domain.UserRetriever
}

func NewAuth(authenticator domain.UserAuthenticator, retriever domain.UserRetriever) (domain.AuthUsecase, error) {
	if retriever == nil {
		return nil, errors.New("user retriever must not be nil")
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
		return "", errors.New("not authorized")
	}

	token, err := u.authenticator.GenerateAccessToken(user)
	if err != nil {
		return "", errors.New("error while generating token")
	}

	return token, nil
}
