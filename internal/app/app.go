package app

import (
	"CRUD/configs"
	"CRUD/internal/database"
	"CRUD/internal/domain"
	"CRUD/internal/router"
	"log"
	"os"
)

func Run() {
	// Env init
	configs.GetEnv()

	// Database init
	database.InitDatabase()
	err := database.DB.AutoMigrate(&domain.Book{})
	if err != nil {
		log.Fatal(err)
	}

	// Router init
	r := router.InitRouter()
	err = r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
