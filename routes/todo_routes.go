package routes

import (
	"todo-api-golang/handlers"
	"todo-api-golang/repositories"
	"todo-api-golang/services"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine) {
	// DEPENDENCIES INJECTION ROUTES
	todoRepository := &repositories.TodoRepository{}
	todoService := services.NewTodoService(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoService)

	// endpoint
	router.GET("/todos", todoHandler.GetTodos)
	router.GET("/todos/:id", todoHandler.GetTodoByID)
	router.POST("/todos", todoHandler.CreateTodo)
	router.PUT("/todos/:id", todoHandler.UpdateTodo)
	router.DELETE("/todos/:id", todoHandler.DeleteTodo)
}
