package handlers

import (
	"Pet-project/internal/taskService"
	"Pet-project/internal/web/tasks"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(service taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// GetTasks обработчик для получения задач
func (h *TaskHandler) GetTasks(ctx echo.Context) error {
	userID := uint(1) // Пример, можно извлечь ID пользователя из контекста
	tasks, err := h.service.GetAllTasks(userID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, tasks)
}

// CreateTask обработчик для создания задачи
func (h *TaskHandler) PostTasks(ctx echo.Context) error {
	var newTask tasks.PostTasksJSONRequestBody
	if err := ctx.Bind(&newTask); err != nil {
		return err
	}
	userID := uint(1) // Пример, можно извлечь ID пользователя из контекста
	task, err := h.service.CreateTask(newTask.Task, userID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, task)
}

// DeleteTasksId обработчик для удаления задачи по ID
func (h *TaskHandler) DeleteTasksId(ctx echo.Context, id int) error {
	taskID := uint(id)
	userID := uint(1) // Пример, можно извлечь ID пользователя из контекста
	if err := h.service.DeleteTask(taskID, userID); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent) // Возвращаем 204 статус
}

// UpdateTask обработчик для обновления задачи
func (h *TaskHandler) PatchTasksId(ctx echo.Context, id int) error {
	var updateData tasks.PatchTasksIdJSONRequestBody
	if err := ctx.Bind(&updateData); err != nil {
		return err
	}
	taskID := uint(id)
	userID := uint(1) // Пример, можно извлечь ID пользователя из контекста
	updatedTask, err := h.service.UpdateTask(taskID, userID, updateData.Task, updateData.IsDone)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, updatedTask)
}
