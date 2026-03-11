package handlers

import (
	"net/http"
	"todo-api-golang/database"
	"todo-api-golang/models"

	"github.com/gin-gonic/gin"
)

// HANDLER FUNCTION GET TODOS
func GetTodos(c *gin.Context) {
	// SLICE ARRAY
	var todos []models.Todo

	// MENEMUKAN SEMUA TODO
	database.DB.Find(&todos)

	// STATUS OK, MENGAMBIL TODO -> MENAMPILKAN JSON
	c.JSON(http.StatusOK, todos)
}

// HANDLER FUNCTION CREATE TODO
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	// Membaca body request JSON
	if err := c.ShouldBindJSON(&todo); err != nil {
		// JIKA ERROR ATAU JSON TIDAK VALID, MAKA AKAN MENGEMBALIKAN ERROR
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// MEMBUAT TODO
	database.DB.Create(&todo)

	// STATUS OK, TERBUAT, DARI TODO BARU
	c.JSON(http.StatusCreated, todo)
}

// HANDLER FUNCTION GET TODOS BY ID
func GetTodoByID(c *gin.Context) {
	// mengambil parameter id dari URL
	id := c.Param("id")

	var todo models.Todo

	// MENGECEK PATH ID DENGAN DATA DI DATABASE
	if err := database.DB.First(&todo, id).Error; err != nil {
		// jika tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// HANDLER FUNCTION UPDATE TODO
func UpdateTodo(c *gin.Context) {

	// ambil id dan deklarasi update todo
	id := c.Param("id")
	var todo models.Todo

	// MENGECEK PATH ID DENGAN DATA DI DATABASE
	if err := database.DB.First(&todo, id).Error; err != nil {
		// jika tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})

		return
	}

	var updateTodo models.Todo
	// MENGECEK VALID DATA UPDATE DI JSON
	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		// JIKA ERROR ATAU JSON TIDAK VALID, MAKA AKAN MENGEMBALIKAN ERROR
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// MENGUPDATE NILAI
	todo.Title = updateTodo.Title
	todo.Completed = updateTodo.Completed

	// MENYIMPAN UPDATE
	database.DB.Save(&todo)

	// STATUS OK
	c.JSON(http.StatusOK, todo)
}

// HANDLER FUNCTION DELETE TODO
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		// jika tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})

		return
	}

	// HAPUS DATA BY ID
	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"message": "todo deleted",
	})
}
