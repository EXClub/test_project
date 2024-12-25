package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordCost int = 14
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) GenerateHash(password []byte) ([]byte, error) {
	if len(string(password)) <= 0 {
		return nil, errors.New("password can ot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword(password, passwordCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (s *AuthService) CompareHash() {
}

func (s *AuthService) GenerateToken() {}

func (s *AuthService) CompareToken() {

}
