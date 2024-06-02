package main

import (
	"go-blog/database"
	"go-blog/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviroment variables
	godotenv.Load()

	// Database
	database.Connect(tools.DBConfig{
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}.String())

	// Router
	router := gin.Default()
	routes.IntializeRoutes(router)
	router.Run(":8080")
}
