package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fakh1m/todoApp/database"
	"github.com/fakh1m/todoApp/models"
	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}
	json.NewEncoder(w).Encode(todos)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID yang kamu masukkan tidak ada, silakan periksa kembali :)", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	err = database.DB.QueryRow("SELECT id, title, completed FROM todos WHERE id =$1", id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err == sql.ErrNoRows {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.DB.QueryRow("INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id", todo.Title, todo.Completed).Scan(&todo.ID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil{
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err:= database.DB.Exec("UPDATE todos SET title=$1, completed=$2 WhHERE id=$3", todo.Title, todo.Completed, id)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func DeleteTodo (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID tidak terdaftar", http.StatusBadRequest)
		return 
	}
	
	result, err := database.DB.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}