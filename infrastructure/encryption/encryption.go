package encryption

import (
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/firdasafridi/gocrypt"
)

var errInternalServerError = errors.New("internal server error")

type encryptor struct {
	des *gocrypt.DESOpt
}

func New(key string) (domain.SummaryEncryptor, error) {
	if key == "" {
		return &encryptor{}, errors.New("key must not be empty")
	}

	des, err := gocrypt.NewDESOpt(key)
	if err != nil {
		return &encryptor{}, err
	}

	return &encryptor{des}, nil
}

func (e *encryptor) Encrypt(value string) (string, error) {
	cipherText, err := e.des.Encrypt([]byte(value))
	if err != nil {
		return "", errInternalServerError
	}

	return cipherText, nil
}

func (e *encryptor) Decrypt(value string) (string, error) {
	plainText, err := e.des.Decrypt([]byte(value))
	if err != nil {
		return "", errInternalServerError
	}

	return plainText, nil
}
