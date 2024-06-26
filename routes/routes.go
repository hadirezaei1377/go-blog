package routes

import (
	"go-blog/controllers"
	"go-blog/middlewares"
	"go-blog/models/permissions"

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

	category_routes := router.Group("api/v1/categories")
	{
		category_routes.POST("/", middlewares.RequireLogin, middlewares.CheckPermissions(permissions.CreateCategory), controllers.CreateCategory)
		category_routes.GET("/:id", controllers.GetCategory)
		category_routes.GET("/", controllers.GetCategories)
	}

	role_routes := router.Group("api/v1/roles")
	{
		// Not implemented routes
		role_routes.POST("/", middlewares.RequireLogin, middlewares.CheckPermissions(permissions.CreateRole), controllers.CreateRole)
		role_routes.PATCH("/", middlewares.RequireLogin, middlewares.CheckPermissions(permissions.UpdateRole), controllers.UpdateRole)
		role_routes.GET("/:id", controllers.GetRole)
		role_routes.GET("/", controllers.GetRoles)
		role_routes.DELETE("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(permissions.DeleteRole), controllers.DeleteRole)
	}
}
