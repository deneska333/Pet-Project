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
	database.InitDB()

	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewService(taskRepo)
	taskHandler := handlers.NewHandler(taskService) // Должен реализовать StrictServerInterface

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewService(userRepo)
	userHandler := handlers.NewUsersHandler(userService) // Аналогично

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasks.RegisterHandlers(e, tasks.NewStrictHandler(taskHandler, nil))
	users.RegisterHandlers(e, users.NewStrictHandler(userHandler, nil))

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
