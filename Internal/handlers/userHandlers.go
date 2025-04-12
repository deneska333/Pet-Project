package handlers

import (
	"Pet-project/internal/userService"
	"Pet-project/internal/web/users"
	"context"
	"errors"

	"gorm.io/gorm"
)

type Handler struct {
	Service *userService.UserService
}

func NewHandler(service *userService.UserService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		id := int64(usr.ID)
		response = append(response, users.User{
			Id:    &id,
			Name:  &usr.Name,
			Email: &usr.Email,
			Role:  (*users.Role)(&usr.Role),
		})
	}
	return response, nil
}

func (h *Handler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userToCreate := userService.User{
		Name:  *request.Body.Name,
		Email: *request.Body.Email,
		Role:  string(*request.Body.Role),
	}

	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	id := int64(createdUser.ID)
	return users.PostUsers201JSONResponse{
		Id:    &id,
		Name:  &createdUser.Name,
		Email: &createdUser.Email,
		Role:  (*users.Role)(&createdUser.Role),
	}, nil
}

func (h *Handler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	updateData := userService.User{}

	if request.Body.Name != nil {
		updateData.Name = *request.Body.Name
	}

	if request.Body.Email != nil {
		updateData.Email = *request.Body.Email
	}

	if request.Body.Role != nil {
		updateData.Role = string(*request.Body.Role)
	}

	updatedUser, err := h.Service.UpdateUser(uint(request.Id), updateData)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return users.PatchUsersId404Response{}, nil
	}
	if err != nil {
		return nil, err
	}

	id := int64(updatedUser.ID)
	return users.PatchUsersId200JSONResponse{
		Id:    &id,
		Name:  &updatedUser.Name,
		Email: &updatedUser.Email,
		Role:  (*users.Role)(&updatedUser.Role),
	}, nil
}

func (h *Handler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.Service.DeleteUser(uint(request.Id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return users.DeleteUsersId404Response{}, nil
	}
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
