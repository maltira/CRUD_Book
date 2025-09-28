package repository

import (
	"CRUD/internal/domain"

	"gorm.io/gorm"
)

// Repository через GORM работает с БД

type BookRepository interface {
	GetAll() ([]domain.Book, error)
	GetByID(id int) (domain.Book, error)
	Create(book domain.Book) error
	Update(id int, book domain.Book) error
	Delete(id int) error
}

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepo{db: db}
}

func (r *BookRepo) GetAll() ([]domain.Book, error) {
	var books []domain.Book                         // Массив всех книг
	if err := r.db.Find(&books).Error; err != nil { // запись в books и проверка на ошибку
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) GetByID(id int) (domain.Book, error) {
	var book domain.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *BookRepo) Create(book domain.Book) error {
	return r.db.Create(&book).Error
}

func (r *BookRepo) Update(id int, book domain.Book) error {
	res := r.db.Model(&domain.Book{}).Where("id = ?", id).Updates(book)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *BookRepo) Delete(id int) error {
	res := r.db.Unscoped().Delete(&domain.Book{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
