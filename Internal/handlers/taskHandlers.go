package handlers

import (
	"context"
	"errors"
	"project/internal/taskService"
	"project/internal/web/tasks"

	"gorm.io/gorm"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		id := int64(tsk.ID)
		response = append(response, tasks.Task{
			Id:     &id,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		})
	}
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskToCreate := taskService.Task{
		Text:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	id := int64(createdTask.ID)
	return tasks.PostTasks201JSONResponse{
		Id:     &id,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	var text string
	var isDone bool

	if request.Body.Task != nil {
		text = *request.Body.Task
	}

	if request.Body.IsDone != nil {
		isDone = *request.Body.IsDone
	}

	updateData := taskService.Task{
		Text:   text,
		IsDone: isDone,
	}

	updatedTask, err := h.Service.UpdateTask(uint(request.Id), updateData)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tasks.PatchTasksId404Response{}, nil
	}
	if err != nil {
		return nil, err
	}

	id := int64(updatedTask.ID)
	return tasks.PatchTasksId200JSONResponse{
		Id:     &id,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTask(uint(request.Id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tasks.DeleteTasksId404Response{}, nil
	}
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}
