package main

import (
	"log"
	"net/http"

	"github.com/fakh1m/todoApp/database"
	"github.com/fakh1m/todoApp/handler"
	"github.com/gorilla/mux"
)

func main() {
    database.ConnectDB()

    r := mux.NewRouter()

    r.HandleFunc("/todos", handler.GetTodos).Methods("GET")
    r.HandleFunc("/todos/{id}", handler.GetTodoById).Methods("GET")
    r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
    r.HandleFunc("/todos/{id}", handler.UpdateTodo).Methods("PUT")
    r.HandleFunc("/todos/{id}", handler.DeleteTodo).Methods("DELETE")

    log.Println("Server running on port 8080")
    http.ListenAndServe(":8080", r)
}
