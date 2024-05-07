package repository

import (
	"errors"
	"github.com/JakubG-git/ToDoApp/pkg/repository/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserInterface interface {
	GetAll() (model.Users, error)
	GetById(id int) (model.User, error)
	GetByUsername(username string) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) bool
	Delete(id int) error
}

type UserRepository struct {
	Db         *gorm.DB
	bcryptCost int
}

func NewUserRepository(db *gorm.DB, bcryptCost int) *UserRepository {
	return &UserRepository{Db: db, bcryptCost: bcryptCost}
}

func (r *UserRepository) GetAll() (model.Users, error) {
	var users model.Users
	err := r.Db.Find(&users).Omit("password_hash").Error
	return users, err
}

func (r *UserRepository) GetById(id int) (model.User, error) {
	var user model.User
	err := r.Db.Preload("ToDos").First(&user, id).Error
	return user, err
}

func (r *UserRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	err := r.Db.Preload("ToDos").Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) Register(user model.User) (model.User, error) {
	err := r.Db.Create(&user).Error
	return user, err
}

func (r *UserRepository) Update(user model.User) (model.User, error) {
	err := r.Db.Save(&user).Error
	return user, err
}

func (r *UserRepository) Login(username, password string) (model.User, error) {
	user, err := r.GetByUsername(username)
	if err != nil {
		return model.User{}, err
	}
	if !r.ComparePassword(password, user.PasswordHash) {
		return model.User{}, ErrInvalidCredentials
	}
	return user, nil
}

func (r *UserRepository) Delete(id int) error {
	err := r.Db.Delete(&model.User{}, id).Error
	return err
}

func (r *UserRepository) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), r.bcryptCost)
	return string(bytes), err
}

func (r *UserRepository) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
