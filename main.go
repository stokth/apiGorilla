package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	Text string `json:"text"`
}

var tasks = make(map[int]Task)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello, %s", tasks[0].Text)
	} else {
		fmt.Fprintln(w, "Поддерживается только GET-запрос")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var task Task
		json.NewDecoder(r.Body).Decode(&task)
		tasks[0] = task
	} else {
		fmt.Fprintln(w, "Поддерживается только POST-запрос")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
