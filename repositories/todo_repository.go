// UNTUK MENGHANDLE EKSEKUSI KE DATABASE
package repositories

import (
	"todo-api-golang/database"
	"todo-api-golang/models"
)

type TodoRepository struct{}

// FUNGSI TAMBAH
func (r *TodoRepository) Create(todo *models.Todo) error {
	result := database.DB.Create(todo)

	// MENGEMBALIKAN BERHASIL ATAU GAGAL
	return result.Error
}

// FUNGSI AMBIL SEMUA
func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	result := database.DB.Find(&todos)

	return todos, result.Error
}

// FUNGSI AMBIL SATU DATA BY ID
func (r *TodoRepository) GetByID(id string) (models.Todo, error) {
	var todo models.Todo

	result := database.DB.First(&todo, id)

	return todo, result.Error
}

// FUNGSI DELETE
func (r *TodoRepository) Delete(id string) error {
	result := database.DB.Delete(&models.Todo{}, id)

	return result.Error
}

// FUNGSI UPDATE
func (r *TodoRepository) Update(id string, updateTodo models.Todo) (models.Todo, error) {
	var todo models.Todo

	result := database.DB.Find(&todo, id)

	if result.Error != nil {
		return todo, result.Error
	}

	todo.Title = updateTodo.Title
	todo.Completed = updateTodo.Completed

	database.DB.Save(&todo)

	return todo, nil
}
