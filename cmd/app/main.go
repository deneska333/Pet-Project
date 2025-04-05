package main

import (
	"Pet-project/internal/database"
	"Pet-project/internal/handlers"
	"Pet-project/internal/taskService"
	"Pet-project/internal/web/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.InitDB()

	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
