package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (Book) TableName() string {
	return "books" // будет использоваться таблица "my_books"
}
