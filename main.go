package main

// user routes
// database
// register
// login and logout

import (
	"go-blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.IntializeRoutes(router)
	router.Run(":8080")
}
