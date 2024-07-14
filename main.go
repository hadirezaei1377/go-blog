package main

import (
	"go-blog/database"
	"go-blog/routes"
	"go-blog/tools"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	godotenv.Load()

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
	routes.InitializeRoutes(router, db)
	router.StaticFS("/", os.DirFS("./ui"))

	router.Start(":80")
}
