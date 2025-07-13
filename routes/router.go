package routes

import (
	"task-service/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(handler *handlers.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	return r
}
