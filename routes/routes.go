package routes

import (
	"go-personal-boilerplate/controllers"
	"go-personal-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	// Protected routes for Todos
	protectedTodos := router.Group("/todos")
	protectedTodos.Use(middlewares.ProtectRoute) // Apply ProtectRoute middleware to all /todos endpoints
	{
		protectedTodos.POST("/", controllers.CreateTodoHandler)
		protectedTodos.GET("/", controllers.GetTodosHandler)
		protectedTodos.PUT("/update/:id", controllers.UpdateTodoHandler)
		protectedTodos.PUT("/status/:id", controllers.UpdateTodoStatusHandler)
		protectedTodos.DELETE("/:id", controllers.DeleteTodoHandler)
	}

	// Protected routes for Posts
	protectedPosts := router.Group("/posts")
	protectedPosts.Use(middlewares.ProtectRoute) // Apply ProtectRoute middleware to all /posts endpoints
	{
		protectedPosts.POST("/", controllers.CreatePostHandler)
		protectedPosts.GET("/", controllers.GetPostsHandler)
		protectedPosts.PUT("/:id", controllers.UpdatePostHandler)
		protectedPosts.DELETE("/:id", controllers.DeletePostHandler)
	}
}
