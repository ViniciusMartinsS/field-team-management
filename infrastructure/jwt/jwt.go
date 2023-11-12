package jwt

import (
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/golang-jwt/jwt/v5"
)

type authenticator struct {
	key string
}

func New(key string) (domain.UserAuthenticator, error) {
	if key == "" {
		return nil, errors.New("key must not be empty")
	}

	return &authenticator{key: key}, nil
}

func (a *authenticator) GenerateAccessToken(user domain.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"role_id": user.RoleID,
		},
	)

	token, err := t.SignedString([]byte(a.key))
	if err != nil {
		return "", err
	}

	return token, nil
}
