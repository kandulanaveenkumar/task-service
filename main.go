package main

import (
	"log"
	"net/http"
	"task-service/routes"
	"task-service/services"
	"task/database"
	"task/handlers"
	"task/repositories"
)

func main() {
	database.ConnectDB()

	taskRepo := repositories.NewTaskRepository(database.DB)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := routes.RegisterRoutes(taskHandler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
