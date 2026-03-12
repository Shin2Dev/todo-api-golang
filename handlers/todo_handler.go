// SEBAGAI CONTROLLER, MENGHUBUNGKAN SERVICE DAN REPOSITORY
package handlers

import (
	"net/http"
	"todo-api-golang/models"
	"todo-api-golang/services"

	"github.com/gin-gonic/gin"
)

// HANDLER MENGGUNAKAN SERVICE
type TodoHandler struct {
	service *services.TodoService
}

// FUNGSI UNTUK MEMBUAT HANDLER BARU DI SETIAP FUNGSI SERVICE API
func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

// HANDLER FUNCTION GET TODOS
func (h *TodoHandler) GetTodos(c *gin.Context) {
	// MENEMUKAN SEMUA TODO
	todos, err := h.service.GetTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get todos",
		})

		return
	}

	// STATUS OK, MENGAMBIL TODO -> MENAMPILKAN JSON
	c.JSON(http.StatusOK, todos)
}

// HANDLER FUNCTION CREATE TODO
func (h *TodoHandler) CreateTodo(c *gin.Context) {
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
	err := h.service.CreateTodo(&todo)

	// JIKA GAGAL BUAT TODO
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create todo",
		})

		return
	}

	// STATUS OK, TERBUAT, DARI TODO BARU
	c.JSON(http.StatusCreated, todo)
}

// HANDLER FUNCTION GET TODOS BY ID
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	// mengambil parameter id dari URL
	id := c.Param("id")

	todo, err := h.service.GetTodoByID(id)

	// MENGECEK PATH ID DENGAN DATA DI DATABASE
	if err != nil {
		// jika tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"error": "todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// HANDLER FUNCTION UPDATE TODO
func (h *TodoHandler) UpdateTodo(c *gin.Context) {

	// ambil id dan deklarasi update todo
	id := c.Param("id")
	var updateTodo models.Todo

	// MENGECEK VALID DATA UPDATE DI JSON
	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		// JIKA ERROR ATAU JSON TIDAK VALID, MAKA AKAN MENGEMBALIKAN ERROR
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.service.GetTodoByID(id)

	// MENGECEK PATH ID DENGAN DATA DI DATABASE
	if err != nil {
		// jika tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"error": "todo not found",
		})

		return
	}

	_, err = h.service.UpdateTodo(id, updateTodo)

	if err != nil {
		// jika tidak ditemukan
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update todo",
		})

		return
	}

	// STATUS OK
	c.JSON(http.StatusOK, updateTodo)
}

// HANDLER FUNCTION DELETE TODO
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	// HAPUS DATA BY ID
	err := h.service.DeleteTodo(id)

	if err != nil {
		// jika tidak ditemukan
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete todo",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo deleted",
	})
}
