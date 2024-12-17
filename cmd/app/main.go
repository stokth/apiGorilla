package main

import (
	"apiGorilla/internal/database"
	"apiGorilla/internal/handlers"
	"apiGorilla/internal/taskService"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Вызываем инициализацию базы данных
	database.InitDB()

	// Автоматическая миграция модели Message
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/posts", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
