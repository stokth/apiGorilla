package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", HelloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
