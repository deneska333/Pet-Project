package handlers

import (
	"log"
	"net/http"
	"strconv"

	"Pet-project/internal/taskService"
	"Pet-project/internal/web/tasks"

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
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get tasks")
	}
	return ctx.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostTasks(ctx echo.Context) error {
	var req tasks.PostTasksJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if req.Task == nil || *req.Task == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Task field is required")
	}

	var userID uint = 1
	if req.UserID != nil {
		userID = uint(*req.UserID)
	}

	task, err := h.service.CreateTask(*req.Task, userID)
	if err != nil {
		if err.Error() == "user not found" {
			return echo.NewHTTPError(http.StatusBadRequest, "Specified user does not exist")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create task")
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
	log.Println(taskID, userID)
	if err := h.service.DeleteTask(taskID, userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete task")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *TaskHandler) PatchTasksId(ctx echo.Context, id int) error {
	var req tasks.PatchTasksIdJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	taskID := uint(id)
	var userID uint = 0

	if req.UserID != nil {
		userID = uint(*req.UserID)
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

	taskText := ""
	if req.Task != nil {
		taskText = *req.Task
	}

	isDone := false
	if req.IsDone != nil {
		isDone = *req.IsDone
	}

	updatedTask, err := h.service.UpdateTask(taskID, userID, taskText, isDone)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update task")
	}

	return ctx.JSON(http.StatusOK, updatedTask)
}
