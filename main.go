package main

import (
	"CRUD/src/config"
	"CRUD/src/controller"
	"CRUD/src/entity"
	"log"

	_ "CRUD/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func init() {
	config.GetEnv()
	config.InitDatabase()

	err := config.DB.AutoMigrate(&entity.Book{})
	if err != nil {
		log.Fatal(err)
	}

}

// @title CRUD_Book API
// @version 1.0
// @description API Server for CRUD_Book Application

// @host localhost:3000
// @BasePath /

func main() {
	r := gin.Default()

	r.GET("/books", controller.BookIndex)
	r.GET("/books/:id", controller.BookById)
	r.POST("/books", controller.BookPost)
	r.PUT("/books/:id", controller.BookPut)
	r.DELETE("/books/:id", controller.BookDelete)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
