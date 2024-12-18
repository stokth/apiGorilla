package handlers

import (
	"apiGorilla/internal/taskService"
	"apiGorilla/internal/web/tasks"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *taskService.TaskService
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskByID(request.Id)

	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(request.Id, taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Неверный тип идентификатора задачи", http.StatusBadRequest)
	}

	var updatedTask taskService.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	updatedTask, err = h.Service.UpdateTaskByID(uint(id), updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Неверный тип идентификатора задачи", http.StatusBadRequest)
	}

	err = h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		http.Error(w, "Не удалось удалить таску", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
}
