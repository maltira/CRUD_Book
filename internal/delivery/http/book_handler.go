package http

import (
	"CRUD/internal/domain"
	"CRUD/internal/service"
	"errors"
	"net/http"
	"strconv"

	_ "CRUD/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// Handler вызывает метод из BookService

type BookResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type BookHandler struct {
	sc service.BookService
}

func NewBookHandler(r *gin.Engine, sc service.BookService) {
	h := &BookHandler{sc: sc}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/books", h.GetAllBooks)
	r.GET("/books/:id", h.GetBookByID)
	r.POST("/books", h.CreateBook)
	r.PUT("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)
}

// @Summary Get list of books
// @Description Возвращает все книги из базы
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} BookResponse
// @Router /books [get]
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.sc.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, books)
}

// @Summary Get book by id
// @Description Возвращает книгу из базы по id
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book_ID"
// @Success 200 {object} BookResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.sc.GetByID(int(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "book not found"})
		return
	}
	c.JSON(200, book)
}

// @Summary Add a new book
// @Description Создаёт новую книгу
// @Tags books
// @Accept json
// @Produce json
// @Param book body domain.Book true "Book info"
// @Router /books/{id} [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book domain.Book
	if err := c.Bind(&book); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if err := h.sc.Create(book); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "created"})
}

// @Summary Update book
// @Description Обновляет книгу по ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body domain.Book true "Updated book info"
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book domain.Book
	if err := c.Bind(&book); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if err := h.sc.Update(int(id), book); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "updated"})
}

// @Summary Delete book
// @Description Удаляет книгу по ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.sc.Delete(int(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "deleted"})
}
