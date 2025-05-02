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
	"gorm.io/gorm"
)

func main() {
	// Инициализация БД
	db := database.InitDB()

	// Создаем временную модель без NOT NULL ограничения
	type TempTask struct {
		gorm.Model
		Text   string `json:"task"`
		IsDone bool   `json:"is_done"`
		UserID uint   `json:"user_id"`
	}

	// 1. Мигрируем временную модель
	if err := db.AutoMigrate(&TempTask{}); err != nil {
		log.Fatalf("failed to auto-migrate temporary task model: %v", err)
	}

	// 2. Проверяем и создаем пользователя по умолчанию если нужно
	var userCount int64
	if err := db.Model(&userService.User{}).Count(&userCount).Error; err != nil {
		log.Fatalf("failed to count users: %v", err)
	}

	if userCount == 0 {
		defaultUser := userService.User{
			Email:    "default@example.com",
			Password: "defaultpassword",
		}
		if err := db.Create(&defaultUser).Error; err != nil {
			log.Fatalf("failed to create default user: %v", err)
		}
	}

	// 3. Обновляем существующие задачи, устанавливая user_id
	if err := db.Model(&TempTask{}).
		Where("user_id IS NULL").
		Update("user_id", 1).Error; err != nil {
		log.Fatalf("failed to update task user_ids: %v", err)
	}

	// 4. Теперь мигрируем настоящую модель
	if err := db.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to auto-migrate final task model: %v", err)
	}

	// Инициализация сервисов и обработчиков
	taskRepo := taskService.NewTaskRepository(db)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := userService.NewUserRepository(db)
	userService := userService.NewService(userRepo)
	userHandler := handlers.NewUsersHandler(userService)

	// Настройка Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация обработчиков
	// Оборачиваем taskHandler и userHandler в интерфейсы
	tasks.RegisterHandlers(e, &tasks.ServerInterfaceWrapper{Handler: taskHandler})
	users.RegisterHandlers(e, &users.ServerInterfaceWrapper{Handler: userHandler})

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
