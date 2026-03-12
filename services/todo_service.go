// MENGURUS SEMUA LOGIC DARI API
package services

import (
	"todo-api-golang/models"
	"todo-api-golang/repositories"
)

// SERVICE MENGGUNAKAN REPOSITORI UNTUK DISETOR KE DATABASE
type TodoService struct {
	repo *repositories.TodoRepository
}

// FUNGSI UNTUK MEMBUAT SERVICE BARU DI SETIAP FUNGSI API
func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

// CREATE
func (s *TodoService) CreateTodo(todo *models.Todo) error {
	return s.repo.Create(todo)
}

// GET ALL TODO
func (s *TodoService) GetTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

// GET ONE TODO
func (s *TodoService) GetTodoByID(id string) (models.Todo, error) {
	return s.repo.GetByID(id)
}

// DELETE
func (s *TodoService) DeleteTodo(id string) error {
	return s.repo.Delete(id)
}

// UPDATE
func (s *TodoService) UpdateTodo(id string, updateTodo models.Todo) (models.Todo, error) {
	return s.repo.Update(id, updateTodo)
}
