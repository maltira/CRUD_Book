package controller

import (
	"CRUD/src/config"
	"CRUD/src/entity"

	"github.com/gin-gonic/gin"
)

// *gin.Context - хранит информацию о запросе и
// позволяет формировать ответ (JSON, HTML, ошибки и т.д.)

// BookResponse swagger-модель для Book
type BookResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// @Summary Get list of books
// @Description Возвращает все книги из базы
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} BookResponse
// @Router /books [get]
func BookIndex(c *gin.Context) {
	var books []entity.Book // Массив всех книг
	config.DB.Find(&books)  // запись в books

	c.JSON(200, gin.H{
		"code":  200,
		"books": books,
	})
}

// @Summary Get book by id
// @Description Возвращает книгу из базы по id
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book_ID"
// @Success 200 {object} BookResponse
// @Router /books/{id} [get]
func BookById(c *gin.Context) {
	id := c.Param("id")
	var book entity.Book

	config.DB.First(&book, id)

	c.JSON(200, gin.H{
		"code": 200,
		"book": book,
	})
}

// @Summary Add a new book
// @Description Создаёт новую книгу
// @Tags books
// @Accept json
// @Produce json
// @Param book body entity.Book true "Book info"
// @Router /books/{id} [post]
func BookPost(c *gin.Context) {
	var book entity.Book
	c.Bind(&book) // берёт данные из тела запроса и заполняет структуру book

	config.DB.Create(&book)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Book added successfully",
	})
}

// @Summary Update book
// @Description Обновляет книгу по ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body entity.Book true "Updated book info"
// @Router /books/{id} [put]
func BookPut(c *gin.Context) {
	id := c.Param("id")
	var book entity.Book

	config.DB.First(&book, id)

	var updatedBook entity.Book
	c.Bind(&updatedBook)

	config.DB.Model(&book).Updates(entity.Book{
		Title:  updatedBook.Title,
		Author: updatedBook.Author,
	})

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Book updated successfully",
	})
}

// @Summary Delete book
// @Description Удаляет книгу по ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Router /books/{id} [delete]
func BookDelete(c *gin.Context) {
	id := c.Param("id")
	config.DB.Unscoped().Delete(&entity.Book{}, id)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Book deleted successfully",
	})
}
