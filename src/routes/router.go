package routes

import (
	"todolist/src/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	return router
}
