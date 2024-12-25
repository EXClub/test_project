package service

import (
	"log"

	"github.com/EXClub/test_project.git/internal/user/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser() error {
	if err := s.repo.SaveUser(); err != nil {
		return err
	}
	log.Println("Пользовтаель успешно добавлен")
	return nil
}
