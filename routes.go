package main

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	{
		admin := r.Group("/admin")
		admin.GET("/stats", statsEndpoint)
		admin.POST("/reset", resetEndpoint)
	}

	{
		user := r.Group("/user")
		user.GET("/tasks", tasksEndpoint)
		user.POST("/task", taskEndpoint)
		user.PUT("/task/:id", updateTaskEndpoint)
		user.DELETE("/task/:id", deleteTaskEndpoint)
	}
}
