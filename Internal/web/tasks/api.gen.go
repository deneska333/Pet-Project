// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Task defines model for Task.
type Task struct {
	Id     *int64  `json:"id,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
	Task   *string `json:"task,omitempty"`
	UserID *int64  `json:"user_id,omitempty"` // Добавлено поле user_id
}

// TaskUpdate defines model for TaskUpdate.
type TaskUpdate struct {
	IsDone *bool   `json:"is_done,omitempty"`
	Task   *string `json:"task,omitempty"`
	UserID *int64  `json:"user_id,omitempty"` // Добавлено поле user_id
}

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody = Task

// PatchTasksIdJSONRequestBody defines body for PatchTasksId for application/json ContentType.
type PatchTasksIdJSONRequestBody = TaskUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx echo.Context) error
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx echo.Context) error
	// Delete a task
	// (DELETE /tasks/{id})
	DeleteTasksId(ctx echo.Context, id int) error
	// Update a task
	// (PATCH /tasks/{id})
	PatchTasksId(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx)
	return err
}

// PostTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PostTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTasks(ctx)
	return err
}

// DeleteTasksId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTasksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTasksId(ctx, id)
	return err
}

// PatchTasksId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchTasksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchTasksId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.POST(baseURL+"/tasks", wrapper.PostTasks)
	router.DELETE(baseURL+"/tasks/:id", wrapper.DeleteTasksId)
	router.PATCH(baseURL+"/tasks/:id", wrapper.PatchTasksId)

}

type GetTasksRequestObject struct {
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type GetTasks200JSONResponse []Task

func (response GetTasks200JSONResponse) VisitGetTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasksRequestObject struct {
	Body *PostTasksJSONRequestBody
}

type PostTasksResponseObject interface {
	VisitPostTasksResponse(w http.ResponseWriter) error
}

type PostTasks201JSONResponse Task

func (response PostTasks201JSONResponse) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTasksIdRequestObject struct {
	Id int `json:"id"`
}

type DeleteTasksIdResponseObject interface {
	VisitDeleteTasksIdResponse(w http.ResponseWriter) error
}

type DeleteTasksId204Response struct {
}

func (response DeleteTasksId204Response) VisitDeleteTasksIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteTasksId404Response struct {
}

func (response DeleteTasksId404Response) VisitDeleteTasksIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PatchTasksIdRequestObject struct {
	Id   int `json:"id"`
	Body *PatchTasksIdJSONRequestBody
}

type PatchTasksIdResponseObject interface {
	VisitPatchTasksIdResponse(w http.ResponseWriter) error
}

type PatchTasksId200JSONResponse Task

func (response PatchTasksId200JSONResponse) VisitPatchTasksIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchTasksId404Response struct {
}

func (response PatchTasksId404Response) VisitPatchTasksIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}
