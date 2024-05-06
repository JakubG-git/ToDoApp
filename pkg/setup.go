package pkg

import (
	"ToDoApp/pkg/auth"
	"ToDoApp/pkg/config"
	"ToDoApp/pkg/controller"
	"ToDoApp/pkg/logging"
	"ToDoApp/pkg/repository"
	"ToDoApp/pkg/repository/model"
	"ToDoApp/pkg/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRoutes(e *gin.Engine, config config.DatabaseConfig) {
	logging.Logger.Info("Setting up routes")
	logging.Logger.Info("Configuring database", zap.String("dsn", config.ParseDSN()))
	gormDb, err := repository.ConfigureDataSource(config.ParseDSN())
	if err != nil {
		panic(err)
	}
	logging.Logger.Info("Migrating database")
	err = repository.Migration(gormDb, &model.User{}, &model.ToDo{})
	if err != nil {
		panic(err)
	}
	setupUserRoutes(e, gormDb)
	setupToDoRoutes(e, gormDb)
	setupHealthCheckRoute(e)

}

func setupUserRoutes(e *gin.Engine, db *gorm.DB) {
	logging.Logger.Info("Setting up user routes")
	userController := controller.NewUserController(service.NewUserService(repository.NewUserRepository(db, 16)))
	usersPublic := e.Group("/users")
	{
		usersPublic.GET("/", userController.GetAll)
		usersPublic.GET("/:id", userController.GetById)
		usersPublic.POST("/register", userController.Register)
		usersPublic.POST("/login", userController.Login)
		usersPublic.POST("/logout", userController.Logout)
	}
	usersPrivate := e.Group("/users")
	{
		usersPrivate.Use(auth.AuthMiddleware())
		usersPrivate.PUT("/:id", userController.Update)
		usersPrivate.DELETE("/:id", userController.Delete)
	}
}

func setupToDoRoutes(e *gin.Engine, db *gorm.DB) {
	logging.Logger.Info("Setting up todo routes")
	todoController := controller.NewToDoController(service.NewToDoService(repository.NewToDoRepository(db)))
	todosPrivate := e.Group("/todos")
	{
		todosPrivate.Use(auth.AuthMiddleware())
		todosPrivate.GET("/", todoController.GetAll)
		todosPrivate.GET("/:id", todoController.GetById)
		todosPrivate.POST("/", todoController.Create)
		todosPrivate.PUT("/:id", todoController.Update)
		todosPrivate.PUT("/:id/complete", todoController.Complete)
		todosPrivate.DELETE("/:id", todoController.Delete)
	}
}

func setupHealthCheckRoute(e *gin.Engine) {
	logging.Logger.Info("Setting up health check route")
	e.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
