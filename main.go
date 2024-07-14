package main

import (
	"go-blog/database"
	"go-blog/routes"
	"go-blog/tools"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	router := gin.Default()
	routes.InitializeRoutes(router, db)
	router.StaticFS("/ui", http.Dir("./ui"))

	router.Run(":80")
}
