package service

import (
	model2 "ToDoApp/pkg/controller/model"
	"ToDoApp/pkg/repository"
	"ToDoApp/pkg/repository/model"
)

type UserServiceInterface interface {
	GetAll() (model.Users, error)
	GetById(id int) (model.User, error)
	GetByUserName(username string) (model.User, error)
	Register(user model2.UserCreateOrUpdateRequest) (model.User, error)
	Update(user model2.UserCreateOrUpdateRequest) (model.User, error)
	Delete(id int) error
	Login(username, password string) (model.User, error)
}

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{UserRepository: ur}
}

func (s *UserService) GetAll() (model.Users, error) {
	return s.UserRepository.GetAll()
}

func (s *UserService) GetById(id int) (model.User, error) {
	return s.UserRepository.GetById(id)
}

func (s *UserService) GetByUserName(username string) (model.User, error) {
	return s.UserRepository.GetByUsername(username)
}

func (s *UserService) Register(user model2.UserCreateOrUpdateRequest) (model.User, error) {
	var userModel model.User
	userModel.Username = user.Username
	passwordHash, err := s.UserRepository.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}
	userModel.PasswordHash = passwordHash
	return s.UserRepository.Register(userModel)
}

func (s *UserService) Update(user model2.UserCreateOrUpdateRequest) (model.User, error) {
	var userModel model.User
	userModel.Username = user.Username
	passwordHash, err := s.UserRepository.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}
	userModel.PasswordHash = passwordHash

	return s.UserRepository.Update(userModel)
}

func (s *UserService) Delete(id int) error {
	return s.UserRepository.Delete(id)
}

func (s *UserService) Login(username, password string) (model.User, error) {
	return s.UserRepository.Login(username, password)
}
