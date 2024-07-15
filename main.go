package main

import (
	database "go-blog/databases"
	"go-blog/routes"
	"go-blog/tools"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	godotenv.Load()

	logger := tools.InitializeLogger()

	// Database
	dsn := tools.DBConfig{
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("HOST"),
	}
	db := database.Connect(dsn.String())

	// Router
	router := echo.New()
	routes.InitializeRoutes(router, db, logger)
	router.StaticFS("/", os.DirFS("./ui"))

	router.Start(":80")
}
