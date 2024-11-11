package encryptutils

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptEncryptor interface {
	Hash(password string) (string, error)
	Check(password, hash string) bool
}

type bcryptEncryptorImpl struct {
	cost int
}

func NewBcryptEncryptor(cost int) *bcryptEncryptorImpl {
	return &bcryptEncryptorImpl{
		cost: cost,
	}
}

func (e *bcryptEncryptorImpl) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), e.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (e *bcryptEncryptorImpl) Check(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
