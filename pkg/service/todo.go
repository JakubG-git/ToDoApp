package service

import (
	model2 "github.com/JakubG-git/ToDoApp/pkg/controller/model"
	"github.com/JakubG-git/ToDoApp/pkg/repository"
	"github.com/JakubG-git/ToDoApp/pkg/repository/model"
)

type ToDoServiceInterface interface {
	GetAll() (model.ToDos, error)
	GetById(id int) (model.ToDo, error)
	Create(todo model2.ToDoCreateOrUpdateRequest, userId uint) (model.ToDo, error)
	Update(todo model2.ToDoCreateOrUpdateRequest, userId, todoId uint) (model.ToDo, error)
	Complete(id int) (model.ToDo, error)
	Delete(id int) error
}

type ToDoService struct {
	ToDoRepository *repository.ToDoRepository
}

func NewToDoService(tr *repository.ToDoRepository) *ToDoService {
	return &ToDoService{ToDoRepository: tr}
}

func (s *ToDoService) GetAll() (model.ToDos, error) {
	return s.ToDoRepository.GetAll()
}

func (s *ToDoService) GetById(id int) (model.ToDo, error) {
	return s.ToDoRepository.GetById(id)
}

func (s *ToDoService) Create(todo model2.ToDoCreateOrUpdateRequest, userId uint) (model.ToDo, error) {
	var todoDb model.ToDo
	todoDb.Title = todo.Title
	todoDb.Description = todo.Description
	todoDb.Done = todo.Done
	todoDb.UserId = userId
	return s.ToDoRepository.Create(todoDb)
}

func (s *ToDoService) Update(todo model2.ToDoCreateOrUpdateRequest, userId, todoId uint) (model.ToDo, error) {
	var todoDb model.ToDo
	todoDb.Title = todo.Title
	todoDb.Description = todo.Description
	todoDb.Done = todo.Done
	todoDb.UserId = userId
	todoDb.ID = todoId
	return s.ToDoRepository.Update(todoDb)
}

func (s *ToDoService) Complete(id int) (model.ToDo, error) {
	return s.ToDoRepository.Complete(id)
}

func (s *ToDoService) Delete(id int) error {
	return s.ToDoRepository.Delete(id)
}
