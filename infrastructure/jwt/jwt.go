package jwt

import (
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/golang-jwt/jwt/v5"
)

type authenticator struct {
	key string
}

func New(key string) (domain.Authenticator, error) {
	if key == "" {
		return &authenticator{}, errors.New("key must not be empty")
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

func (a *authenticator) IsAccessTokenValid(token string) (bool, map[string]any, error) {
	claims := make(jwt.MapClaims)

	parsedToken, err := jwt.ParseWithClaims(token, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(a.key), nil
		},
	)
	if err != nil {
		return false, nil, err
	}

	claimsResult := make(map[string]any)
	for k, c := range claims {
		claimsResult[k] = c
	}

	return parsedToken.Valid, claimsResult, err
}
