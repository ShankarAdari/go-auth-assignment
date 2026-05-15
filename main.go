package main

import (
	"go-auth-assignment/database"
	"go-auth-assignment/handlers"
	"go-auth-assignment/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	// CORS middleware — allow browser frontend to call API
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Serve the frontend UI
	r.Static("/ui", "./static")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Public routes
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", handlers.GetProfile)

		// Admin only route
		admin := auth.Group("/")
		admin.Use(middleware.RoleMiddleware("admin"))
		{
			admin.GET("/users", handlers.GetAllUsers)
		}
	}

	r.Run(":8080")
}
