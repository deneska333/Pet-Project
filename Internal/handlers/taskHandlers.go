package handlers

import (
	"Pet-project/internal/taskService"
	"Pet-project/internal/web/tasks"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(service taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(ctx echo.Context) error {

	userIDStr := ctx.QueryParam("user_id")
	var userID uint = 0

	if userIDStr != "" {
		id, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id parameter")
		}
		userID = uint(id)
	}

	tasks, err := h.service.GetAllTasks(userID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostTasks(ctx echo.Context) error {
	var newTask tasks.PostTasksJSONRequestBody
	if err := ctx.Bind(&newTask); err != nil {
		return err
	}

	var userID uint = 1
	if newTask.UserID != nil {
		userID = uint(*newTask.UserID)
	}

	task, err := h.service.CreateTask(newTask.Task, userID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) DeleteTasksId(ctx echo.Context, id int) error {
	taskID := uint(id)

	userIDStr := ctx.QueryParam("user_id")
	var userID uint = 0

	if userIDStr != "" {
		id, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id parameter")
		}
		userID = uint(id)
	}

	if err := h.service.DeleteTask(taskID, userID); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h *TaskHandler) PatchTasksId(ctx echo.Context, id int) error {
	var updateData tasks.PatchTasksIdJSONRequestBody
	if err := ctx.Bind(&updateData); err != nil {
		return err
	}

	taskID := uint(id)

	var userID uint = 0

	if updateData.UserID != nil {
		userID = uint(*updateData.UserID)
	} else {
		userIDStr := ctx.QueryParam("user_id")
		if userIDStr != "" {
			id, err := strconv.ParseUint(userIDStr, 10, 32)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id parameter")
			}
			userID = uint(id)
		}
	}

	updatedTask, err := h.service.UpdateTask(taskID, userID, *updateData.Task, *updateData.IsDone)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, updatedTask)
}
