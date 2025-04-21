package handlers

import (
	"Pet-project/internal/userService"
	"Pet-project/internal/web/users"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UsersHandler struct {
	Service *userService.UserService
}

func NewUsersHandler(service *userService.UserService) *UsersHandler {
	return &UsersHandler{Service: service}
}

func (h *UsersHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		id := int64(usr.ID)
		response = append(response, users.User{
			Id:    &id,
			Email: &usr.Email,
		})
	}
	return response, nil
}

func (h *UsersHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body.Email == nil {
		return nil, errors.New("email is required")
	}

	userToCreate := userService.User{
		Email:    *request.Body.Email,
		Password: "default_password", // Замените на реальную логику
	}

	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	id := int64(createdUser.ID)
	return users.PostUsers201JSONResponse{
		Id:    &id,
		Email: &createdUser.Email,
	}, nil
}

func (h *UsersHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	updateData := userService.User{}
	if request.Body.Email != nil {
		updateData.Email = *request.Body.Email
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
		Email: &updatedUser.Email,
	}, nil
}

func (h *UsersHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.Service.DeleteUser(uint(request.Id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return users.DeleteUsersId404Response{}, nil
	}
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (h *UsersHandler) GetUsersId(_ context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	user, err := h.Service.GetUserByID(uint(request.Id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return users.GetUsersId404Response{}, nil
	}
	if err != nil {
		return nil, err
	}

	id := int64(user.ID)
	return users.GetUsersId200JSONResponse{
		Id:    &id,
		Email: &user.Email,
	}, nil
}
