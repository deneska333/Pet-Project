package main

import (
	"Pet-project/internal/database"
	"Pet-project/internal/handlers"
	"Pet-project/internal/taskService"
	"Pet-project/internal/userService"

	tasksapi "Pet-project/internal/web/tasks"
	usersapi "Pet-project/internal/web/users"

	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	db := database.InitDB()

	taskRepo := taskService.NewTaskRepository(db)

	taskSvc := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskSvc)

	userRepo := userService.NewUserRepository(db)
	userSvc := userService.NewService(userRepo)
	userHandler := handlers.NewUsersHandler(userSvc)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasksapi.RegisterHandlers(e, taskHandler)
	usersapi.RegisterHandlers(e, userHandler)

	log.Println("⇨ http server started on [::]:8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
