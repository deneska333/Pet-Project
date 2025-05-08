package main

import (
	"Pet-project/internal/database"
	"Pet-project/internal/handlers"
	"Pet-project/internal/taskService"
	"Pet-project/internal/userService"
	"Pet-project/internal/web/tasks"
	"Pet-project/internal/web/users"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db := database.InitDB()

	taskRepo := taskService.NewTaskRepository(db)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := userService.NewUserRepository(db)
	userService := userService.NewService(userRepo)
	userHandler := handlers.NewUsersHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasks.RegisterHandlers(e, taskHandler)
	users.RegisterHandlers(e, userHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
