package repository

import (
	"ToDoApp/pkg/repository/model"
	"gorm.io/gorm"
)

type ToDoInterface interface {
	GetAll() (model.ToDos, error)
	GetById(id int) (model.ToDo, error)
	Create(todo model.ToDo) (model.ToDo, error)
	Update(todo model.ToDo) (model.ToDo, error)
	Complete(id int) (model.ToDo, error)
	Delete(id int) error
}

type ToDoRepository struct {
	Db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) *ToDoRepository {
	return &ToDoRepository{Db: db}
}

func (r *ToDoRepository) GetAll() (model.ToDos, error) {
	var todos model.ToDos
	err := r.Db.Find(&todos).Error
	return todos, err
}

func (r *ToDoRepository) GetById(id int) (model.ToDo, error) {
	var todo model.ToDo
	err := r.Db.First(&todo, id).Error
	return todo, err
}

func (r *ToDoRepository) Create(todo model.ToDo) (model.ToDo, error) {
	err := r.Db.Create(&todo).Error
	return todo, err
}

func (r *ToDoRepository) Update(todo model.ToDo) (model.ToDo, error) {
	err := r.Db.Save(&todo).Error
	return todo, err
}

func (r *ToDoRepository) Complete(id int) (model.ToDo, error) {
	var todo model.ToDo
	err := r.Db.Model(&todo).Where("id = ?", id).Update("done", true).Error
	return todo, err
}

func (r *ToDoRepository) Delete(id int) error {
	var todo model.ToDo
	err := r.Db.Model(&todo).Where("id = ?", id).Delete(&todo).Error
	return err
}
