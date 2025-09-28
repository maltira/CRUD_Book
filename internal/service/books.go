package service

import (
	"CRUD/internal/domain"
	"CRUD/internal/repository"
	"errors"
)

// service проверяет правила и вызывает BookRepository

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(r repository.BookRepository) *BookService {
	return &BookService{repo: r}
}

func (s *BookService) GetAll() ([]domain.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetByID(id int) (domain.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) Create(book domain.Book) error {
	if book.Title == "" || book.Author == "" {
		return errors.New("title and author cannot be empty")
	}
	return s.repo.Create(domain.Book{
		Title:  book.Title,
		Author: book.Author,
	})
}

func (s *BookService) Update(id int, book domain.Book) error {
	return s.repo.Update(id, domain.Book{
		Title:  book.Title,
		Author: book.Author,
	})
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}
