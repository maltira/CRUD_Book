package app

import (
	"CRUD/configs"
	"CRUD/internal/database"
	"CRUD/internal/delivery/http"
	"CRUD/internal/domain"
	"CRUD/internal/repository"
	"CRUD/internal/service"
	"os"

	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	// Env init
	configs.GetEnv()

	// Database init
	db := database.InitDatabase()
	err := db.AutoMigrate(&domain.Book{})
	if err != nil {
		log.Fatal(err)
	}

	// связываем слои
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)

	r := gin.Default()
	http.NewBookHandler(r, *bookService)

	r.Run(":" + os.Getenv("PORT"))
}
