package routes

import (
	"todo-api-golang/handlers"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine) {
	// endpoint pertama
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "server berjalan",
	// 	})
	// })

	// endpoint
	router.GET("/todos", handlers.GetTodos)
	router.GET("/todos/:id", handlers.GetTodoByID)
	router.POST("/todos", handlers.CreateTodo)
	router.PUT("/todos/:id", handlers.UpdateTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)
}
