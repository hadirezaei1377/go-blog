package routes

import (
	"go-blog/controllers"
	"go-blog/middlewares"

	"github.com/gin-gonic/gin"
)

func IntializeRoutes(router *gin.Engine) {
	router.Use(middlewares.CORSMiddleware())
	user_routes := router.Group("/api/v1/users")
	{
		user_routes.POST("/register", controllers.UserRegister)
		user_routes.POST("/login", controllers.UserLogin)
		user_routes.POST("/logout", middlewares.RequireLogin, controllers.UserLogout)
		user_routes.GET("/check_username", controllers.CheckUsername)
		user_routes.GET("/id", middlewares.RequireLogin, controllers.UserID)
		// Get & Update Profile and more
		// user_routes.GET("/info/:id", controllers.UserInfo)
	}

	post_routes := router.Group("api/v1/posts")
	{
		post_routes.POST("/", middlewares.RequireLogin, controllers.CreatePost)
		post_routes.GET("/:id", controllers.GetPost)
		post_routes.GET("/", controllers.GetPosts)
	}
}
