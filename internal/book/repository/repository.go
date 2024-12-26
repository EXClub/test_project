package repository

import "github.com/EXClub/test_project.git/internal/book/model"

type BookRepository interface {
	SaveBook()
	LoadBook()
	GetBook() model.Book
}
