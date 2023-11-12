package jwt

import (
	"errors"
	"fmt"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/golang-jwt/jwt/v5"
)

type authenticator struct {
	key string
}

func New(key string) (domain.Authenticator, error) {
	if key == "" {
		return nil, errors.New("key must not be empty")
	}

	return &authenticator{key: key}, nil
}

func (a *authenticator) GenerateAccessToken(user domain.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": fmt.Sprintf("%d", user.ID),
			"email":   user.Email,
			"role_id": fmt.Sprintf("%d", user.RoleID),
		},
	)

	token, err := t.SignedString([]byte(a.key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *authenticator) IsAccessTokenValid(token string) (bool, map[string]any, error) {
	claims := make(jwt.MapClaims)

	parsedToken, err := jwt.ParseWithClaims(token, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(a.key), nil
		},
	)

	claimsResult := make(map[string]any)
	for k, c := range claims {
		claimsResult[k] = c
	}

	return parsedToken.Valid, claimsResult, err
}
