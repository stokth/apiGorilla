package main

import (
	"apiGorilla/internal/database"
	"apiGorilla/internal/handlers"
	"apiGorilla/internal/taskService"
	"apiGorilla/internal/userService"
	"apiGorilla/internal/web/tasks"
	"apiGorilla/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Вызываем инициализацию базы данных
	database.InitDB()
	// Автоматическая миграция модели Message
	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal(err)
	}

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewService(taskRepo)

	userRepo := userService.NewUsersRepository(database.DB)
	userService := userService.NewService(userRepo)

	TaskHandlers := handlers.NewHandler(taskService)
	UserHandlers := handlers.NewHandlerUser(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(TaskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	strictHandlerUser := users.NewStrictHandler(UserHandlers, nil)
	users.RegisterHandlers(e, strictHandlerUser)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
