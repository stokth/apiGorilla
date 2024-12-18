package main

import (
	"apiGorilla/internal/database"
	"apiGorilla/internal/handlers"
	"apiGorilla/internal/taskService"
	"apiGorilla/internal/web/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Вызываем инициализацию базы данных
	database.InitDB()
	// Автоматическая миграция модели Message
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
