package controller

import (
	"ToDoApp/pkg/controller/model"
	"ToDoApp/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ToDoControllerInterface interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Complete(c *gin.Context)
	Delete(c *gin.Context)
	checkIfUserOwnsTodo(userid uint, todoId int) bool
}

type ToDoController struct {
	ToDoService *service.ToDoService
}

func NewToDoController(ts *service.ToDoService) *ToDoController {
	return &ToDoController{ToDoService: ts}
}

func (tc *ToDoController) GetAll(c *gin.Context) {
	todos, err := tc.ToDoService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (tc *ToDoController) GetById(c *gin.Context) {
	todoId := c.Param("id")
	convId, err := strconv.Atoi(todoId)
	todo, err := tc.ToDoService.GetById(convId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (tc *ToDoController) Create(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found in context please login"})
		return
	}
	var todo model.ToDoCreateOrUpdateRequest
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdToDo, err := tc.ToDoService.Create(todo, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdToDo)
}

func (tc *ToDoController) Update(c *gin.Context) {
	todoId := c.Param("id")
	convId, err := strconv.Atoi(todoId)
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found in context please login"})
		return
	}
	if !tc.checkIfUserOwnsTodo(userId.(uint), convId) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you do not own this todo"})
		return
	}
	var todo model.ToDoCreateOrUpdateRequest
	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedToDo, err := tc.ToDoService.Update(todo, userId.(uint), uint(convId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedToDo)

}

func (tc *ToDoController) Complete(c *gin.Context) {
	todoId := c.Param("id")
	convId, err := strconv.Atoi(todoId)
	todo, err := tc.ToDoService.Complete(convId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (tc *ToDoController) Delete(c *gin.Context) {
	todoId := c.Param("id")
	convId, err := strconv.Atoi(todoId)
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found in context please login"})
		return
	}
	if !tc.checkIfUserOwnsTodo(userId.(uint), convId) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you do not own this todo"})
		return
	}
	err = tc.ToDoService.Delete(convId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func (tc *ToDoController) checkIfUserOwnsTodo(userid uint, todoId int) bool {
	todo, err := tc.ToDoService.GetById(todoId)
	if err != nil {
		return false
	}
	return todo.UserId == userid
}
