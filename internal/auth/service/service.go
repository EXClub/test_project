package service

import (
	"errors"

	"github.com/EXClub/test_project.git/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordCost int = 14
)

type AuthService struct {
	userRepo repository.UserRepository
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

func (s *AuthService) ValidateUser(username string, password []byte) (bool, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(password, passwordCost)
	if err != nil {
		return false, nil
	}
	validHash := s.userRepo.GetHash()

	if err = bcrypt.CompareHashAndPassword(passwordHash, validHash); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (s *AuthService) GenerateToken() {}

func (s *AuthService) CompareToken() {

}
