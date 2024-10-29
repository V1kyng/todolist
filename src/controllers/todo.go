package controllers

import (
	"encoding/json"
	"net/http"
	"todolist/src/models"
)

var todo = []models.Todo{
	{Id: 1, Title: "First row", Completed: true},
	{Id: 2, Title: "Second row", Completed: true},
	{Id: 3, Title: "Third row", Completed: false},
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
