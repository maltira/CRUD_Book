package router

import (
	"CRUD/internal/controller"

	_ "CRUD/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/books", controller.BookIndex)
	r.GET("/books/:id", controller.BookById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/books", controller.BookPost)

	r.PUT("/books/:id", controller.BookPut)

	r.DELETE("/books/:id", controller.BookDelete)

	return r
}
