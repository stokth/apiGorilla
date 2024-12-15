package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func main() {
	// Вызываем инициализацию базы данных
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	router.HandleFunc("/api/posts", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
