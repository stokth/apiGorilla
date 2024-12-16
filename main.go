package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var messages []Message

		if err := DB.Find(&messages).Error; err != nil {
			fmt.Fprintln(w, "Ошибка при получении сообщений")
		}

		json.NewEncoder(w).Encode(&messages)
	} else {
		fmt.Fprintln(w, "Поддерживается только GET-запрос")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var message Message

		json.NewDecoder(r.Body).Decode(&message)

		// Сохраняем сообщение в базе данных
		if err := DB.Create(&message).Error; err != nil {
			fmt.Fprintln(w, "Ошибка при сохранении сообщения")
		}
	} else {
		fmt.Fprintln(w, "Поддерживается только POST-запрос")
	}
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		idParam := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if err != nil {
			fmt.Fprintln(w, "Неверный идентификатор сообщения")
		}

		var updateMessage Message
		var message Message
		json.NewDecoder(r.Body).Decode(&updateMessage)

		if err := DB.First(&message, id).Error; err != nil {
			fmt.Fprintln(w, "Сообщение с таким идентификатором не найдено")
		}

		if updateMessage.Task == "" {
			updateMessage.Task = message.Task
		} else if !updateMessage.IsDone && message.IsDone != updateMessage.IsDone {
			updateMessage.IsDone = message.IsDone
		}

		if err := DB.Model(&Message{}).Where("id = ?", id).Update("is_done", updateMessage.IsDone).Update("task", updateMessage.Task).Error; err != nil {
			fmt.Fprintln(w, "Ошибка при изменении сообщения")
		}
	} else {
		fmt.Fprintln(w, "Поддерживается только PATCH-запрос")
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		idParam := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if err != nil {
			fmt.Fprintln(w, "Неверный идентификатор сообщения")
		}

		if err := DB.Delete(&Message{}, id).Error; err != nil {
			fmt.Fprintln(w, "Ошибка при удалении сообщения")
		}
	} else {
		fmt.Fprintln(w, "Поддерживается только Delete-запрос")
	}
}

func main() {
	// Вызываем инициализацию базы данных
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	router.HandleFunc("/api/posts", PostHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
