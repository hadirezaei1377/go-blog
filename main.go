package main

import (
	database "go-blog/databases/gorm"
	"go-blog/log"

	"go-blog/routes"
	"go-blog/tools"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	logger := log.InitializeLogger()

	if err := godotenv.Load(); err != nil {
		logger.Error(err.Error())
	}

	// Database
	dsn := tools.DBConfig{
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("HOST"),
	}
	db, err := database.Connect(dsn.String())
	if err != nil {
		log.Gl.Error(err.Error())
		return
	}

	// Router
	router := echo.New()
	routes.InitializeRoutes(router, db, logger)
	router.StaticFS("/", os.DirFS("./ui"))

	if err := router.Start(":8080"); err != nil {
		log.Gl.Error(err.Error())
	}
}
