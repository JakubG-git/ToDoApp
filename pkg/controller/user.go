package controller

import (
	"github.com/JakubG-git/ToDoApp/pkg/auth"
	"github.com/JakubG-git/ToDoApp/pkg/controller/model"
	"github.com/JakubG-git/ToDoApp/pkg/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type UserController struct {
	UserService *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{UserService: us}
}

func (us *UserController) Register(c *gin.Context) {
	var user model.UserCreateOrUpdateRequest
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userModel, err := us.UserService.Register(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = auth.GenerateTokenPair(c, userModel.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user registered successfully"})

}

func (us *UserController) Login(c *gin.Context) {
	var user model.UserLoginRequest
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userModel, err := us.UserService.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = auth.GenerateTokenPair(c, userModel.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user logged in successfully"})
}

func (us *UserController) Logout(c *gin.Context) {
	auth.ClearCookies(c)
}

func (us *UserController) GetAll(c *gin.Context) {
	users, err := us.UserService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (us *UserController) GetById(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(400, gin.H{"error": "user id not found in context please login"})
		return
	}
	userIdParam := c.Param("id")
	if userIdParam != userId {
		c.JSON(403, gin.H{"error": "you do not have permission to access this resource"})
		return
	}
	user, err := us.UserService.GetById(userId.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)

}

func (us *UserController) Update(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(400, gin.H{"error": "user id not found in context please login"})
		return
	}
	userIdParam := c.Param("id")
	if userIdParam != userId {
		c.JSON(403, gin.H{"error": "you do not have permission to access this resource"})
		return
	}
	var user model.UserCreateOrUpdateRequest
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := us.UserService.Update(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedUser)

}

func (us *UserController) Delete(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(400, gin.H{"error": "user id not found in context please login"})
		return
	}
	userIdParam := c.Param("id")
	if userIdParam != userId {
		c.JSON(403, gin.H{"error": "you do not have permission to access this resource"})
		return
	}
	err := us.UserService.Delete(userId.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user deleted successfully"})

}
