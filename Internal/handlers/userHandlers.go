package handlers

import (
	"Pet-project/internal/userService"
	"Pet-project/internal/web/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	service userService.UserService
}

func NewUsersHandler(service userService.UserService) *UsersHandler {
	return &UsersHandler{service: service}
}

// GetUsers обработчик для получения пользователей
func (h *UsersHandler) GetUsers(ctx echo.Context) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, users)
}

// CreateUser обработчик для создания пользователя
func (h *UsersHandler) PostUsers(ctx echo.Context) error {
	var newUser users.PostUsersJSONRequestBody
	if err := ctx.Bind(&newUser); err != nil {
		return err
	}
	user := userService.User{
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, createdUser)
}

// DeleteUsersId обработчик для удаления пользователя по ID
func (h *UsersHandler) DeleteUsersId(ctx echo.Context, id int) error {
	userID := uint(id)
	if err := h.service.DeleteUser(userID); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent) // Возвращаем 204 статус
}

// UpdateUser обработчик для обновления пользователя
func (h *UsersHandler) PatchUsersId(ctx echo.Context, id int) error {
	var updateData users.PatchUsersIdJSONRequestBody
	if err := ctx.Bind(&updateData); err != nil {
		return err
	}
	userID := uint(id)
	user := userService.User{
		Email:    updateData.Email,
		Password: updateData.Password,
	}
	updatedUser, err := h.service.UpdateUser(userID, user)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, updatedUser)
}
