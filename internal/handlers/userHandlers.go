package handlers

import (
	"apiGorilla/internal/userService"
	"apiGorilla/internal/web/users"
	"context"
)

type HandlerUsers struct {
	Service *userService.UserService
}

// GetUsersTasksId implements users.StrictServerInterface.
func (h *HandlerUsers) GetUsersTasksId(ctx context.Context, request users.GetUsersTasksIdRequestObject) (users.GetUsersTasksIdResponseObject, error) {
	allTasksForUser, err := h.Service.GetTasksForUser(request.Id)
	if err != nil {
		return nil, err
	}

	response := users.GetUsersTasksId200JSONResponse{}

	for _, tsk := range allTasksForUser {
		tasks := users.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, tasks)
	}

	return response, nil
}

// DeleteUsersId implements users.StrictServerInterface.
func (h *HandlerUsers) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.Service.DeleteUserByID(request.Id)

	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}

// GetUsers implements users.StrictServerInterface.
func (h *HandlerUsers) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.Users{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (h *HandlerUsers) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.Users{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	updatedUser, err := h.Service.UpdateUserByID(request.Id, userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &request.Id,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

// PostUsers implements users.StrictServerInterface.
func (h *HandlerUsers) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.Users{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandlerUser(service *userService.UserService) *HandlerUsers {
	return &HandlerUsers{
		Service: service,
	}
}
