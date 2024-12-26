package service

import "github.com/EXClub/test_project.git/internal/book/model"

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) GetBook() model.Book {
	return model.Book{}
}
