package handlers

import (
	"Pet-project/internal/userService"
	"Pet-project/internal/web/users"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	service userService.UserService
}

func NewUsersHandler(service userService.UserService) *UsersHandler {

	if service == nil {
		log.Fatal("userService cannot be nil in NewUsersHandler")
	}
	return &UsersHandler{service: service}
}

func jsonError(ctx echo.Context, code int, message string) error {
	errResp := users.Error{
		Code:    int32(code),
		Message: message,
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return ctx.JSON(code, errResp)
}

func (h *UsersHandler) GetUsers(ctx echo.Context) error {

	if h.service == nil {
		return jsonError(ctx, http.StatusInternalServerError, "User service is not initialized")
	}

	usersData, err := h.service.GetAllUsers()
	if err != nil {
		log.Printf("Error getting all users: %v", err)
		return jsonError(ctx, http.StatusInternalServerError, "Failed to retrieve users")
	}

	responseUsers := make([]users.User, 0, len(usersData))
	for _, u := range usersData {
		apiUser := users.User{
			Id:    &u.ID,
			Email: u.Email,
		}
		if u.Name != "" {
			apiUser.Name = &u.Name
		}
		if u.Role != "" {
			roleEnum := users.Role(u.Role)
			isValidRole := false
			for _, validRole := range []users.Role{users.Admin, users.UserRole} {
				if roleEnum == validRole {
					isValidRole = true
					break
				}
			}
			if isValidRole {
				apiUser.Role = &roleEnum
			} else {
				log.Printf("Warning: User ID %d has invalid role '%s' from service", u.ID, u.Role)
			}
		}
		responseUsers = append(responseUsers, apiUser)
	}

	return ctx.JSON(http.StatusOK, responseUsers)
}

func (h *UsersHandler) PostUsers(ctx echo.Context) error {
	if h.service == nil {
		return jsonError(ctx, http.StatusInternalServerError, "User service is not initialized")
	}

	var newUserRequest users.PostUsersJSONRequestBody
	if err := ctx.Bind(&newUserRequest); err != nil {
		return jsonError(ctx, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
	}

	if newUserRequest.Email == "" || newUserRequest.Password == "" {
		return jsonError(ctx, http.StatusBadRequest, "Email and password are required")
	}

	userToCreate := userService.User{
		Email:    newUserRequest.Email,
		Password: newUserRequest.Password,
	}
	if newUserRequest.Name != nil {
		userToCreate.Name = *newUserRequest.Name
	}
	if newUserRequest.Role != nil {
		userToCreate.Role = string(*newUserRequest.Role)
	}

	createdUser, err := h.service.CreateUser(userToCreate)
	if err != nil {
		var conflictErr *userService.ErrEmailConflict
		if errors.As(err, &conflictErr) {
			return jsonError(ctx, http.StatusConflict, err.Error())
		}
		log.Printf("Error creating user: %v", err)
		return jsonError(ctx, http.StatusInternalServerError, "Failed to create user")
	}

	responseUser := users.User{
		Id:    &createdUser.ID,
		Email: createdUser.Email,
	}
	if createdUser.Name != "" {
		responseUser.Name = &createdUser.Name
	}
	if createdUser.Role != "" {
		roleEnum := users.Role(createdUser.Role)
		isValidRole := false
		for _, validRole := range []users.Role{users.Admin, users.UserRole} {
			if roleEnum == validRole {
				isValidRole = true
				break
			}
		}
		if isValidRole {
			responseUser.Role = &roleEnum
		}
	}

	return ctx.JSON(http.StatusCreated, responseUser)
}

func (h *UsersHandler) DeleteUsersId(ctx echo.Context, id int64) error {
	if h.service == nil {
		return jsonError(ctx, http.StatusInternalServerError, "User service is not initialized")
	}

	if id <= 0 {
		return jsonError(ctx, http.StatusBadRequest, "Invalid user ID provided")
	}
	userID := uint(id)

	err := h.service.DeleteUser(userID)
	if err != nil {

		var notFoundErr *userService.ErrUserNotFound

		if errors.As(err, &notFoundErr) {

			return ctx.NoContent(http.StatusNotFound)
		}

		log.Printf("Error deleting user %d: %v", id, err)
		return jsonError(ctx, http.StatusInternalServerError, "Failed to delete user")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *UsersHandler) GetUsersId(ctx echo.Context, id int64) error {
	if h.service == nil {
		return jsonError(ctx, http.StatusInternalServerError, "User service is not initialized")
	}

	if id <= 0 {
		return jsonError(ctx, http.StatusBadRequest, "Invalid user ID provided")
	}
	userID := uint(id)

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		var notFoundErr *userService.ErrUserNotFound

		if errors.As(err, &notFoundErr) {
			return jsonError(ctx, http.StatusNotFound, err.Error())
		}

		log.Printf("Error getting user %d: %v", id, err)
		return jsonError(ctx, http.StatusInternalServerError, "Failed to retrieve user")
	}

	responseUser := users.User{
		Id:    &user.ID,
		Email: user.Email,
	}
	if user.Name != "" {
		responseUser.Name = &user.Name
	}
	if user.Role != "" {
		roleEnum := users.Role(user.Role)
		isValidRole := false
		for _, validRole := range []users.Role{users.Admin, users.UserRole} {
			if roleEnum == validRole {
				isValidRole = true
				break
			}
		}
		if isValidRole {
			responseUser.Role = &roleEnum
		}
	}

	return ctx.JSON(http.StatusOK, responseUser)
}

func (h *UsersHandler) PatchUsersId(ctx echo.Context, id int64) error {
	if h.service == nil {
		return jsonError(ctx, http.StatusInternalServerError, "User service is not initialized")
	}

	if id <= 0 {
		return jsonError(ctx, http.StatusBadRequest, "Invalid user ID provided")
	}

	var updateData users.PatchUsersIdJSONRequestBody
	if err := ctx.Bind(&updateData); err != nil {
		return jsonError(ctx, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
	}

	if updateData.Name == nil && updateData.Email == nil && updateData.Role == nil {
		return jsonError(ctx, http.StatusBadRequest, "PATCH request body must contain at least one field to update")
	}

	userID := uint(id)

	updatePayload := userService.UserUpdatePayload{
		Name:  updateData.Name,
		Email: updateData.Email,
	}
	if updateData.Role != nil {
		roleStr := string(*updateData.Role)
		updatePayload.Role = &roleStr
	}

	updatedUser, err := h.service.UpdateUser(userID, updatePayload)
	if err != nil {
		var notFoundErr *userService.ErrUserNotFound
		var conflictErr *userService.ErrEmailConflict

		if errors.As(err, &notFoundErr) {
			return jsonError(ctx, http.StatusNotFound, err.Error())
		}
		if errors.As(err, &conflictErr) {
			return jsonError(ctx, http.StatusConflict, err.Error())
		}
		log.Printf("Error updating user %d: %v", id, err)
		return jsonError(ctx, http.StatusInternalServerError, "Failed to update user")
	}

	responseUser := users.User{
		Id:    &updatedUser.ID,
		Email: updatedUser.Email,
	}
	if updatedUser.Name != "" {
		responseUser.Name = &updatedUser.Name
	}
	if updatedUser.Role != "" {
		roleEnum := users.Role(updatedUser.Role)
		isValidRole := false
		for _, validRole := range []users.Role{users.Admin, users.UserRole} {
			if roleEnum == validRole {
				isValidRole = true
				break
			}
		}
		if isValidRole {
			responseUser.Role = &roleEnum
		}
	}

	return ctx.JSON(http.StatusOK, responseUser)
}
